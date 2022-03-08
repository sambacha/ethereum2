package crawler

import (
	"crypto/ecdsa"
	"log"
	"net"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p/discover"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ppopth/discv5-crawl/param"
)

// Crawler is a container for states of a cralwer node.
type Crawler struct {
	// The private key used to run the ethereum node.
	privateKey *ecdsa.PrivateKey
	// The list of ethereum nodes used to bootstrap the network.
	bootstrapNodes []*enode.Node
	// The interface used to communicate with the ethereum DHT.
	disc *discover.UDPv5
}

// New creates a new crawler.
func New(bootUrls []string) *Crawler {
	// Generate a new key pair every time we create a new crawler.
	// We can ignore an error over here, because it's just a key generation
	// and there will be no error.
	key, _ := crypto.GenerateKey()
	// Parse the bootstrap nodes.
	if bootUrls == nil {
		log.Println("new crawler created with default bootstrap nodes")
		bootUrls = param.V5Bootnodes
	} else {
		log.Println("new crawler created with custom bootstrap nodes")
	}
	bootNodes := parseBoostrapUrls(bootUrls)

	return &Crawler{
		privateKey:     key,
		bootstrapNodes: bootNodes,
	}
}

// Start crawling.
func (c *Crawler) Run() error {
	if err := c.setupDiscovery(); err != nil {
		return err
	}
	// c.disc is produced after setupDiscovery. We need to eventually close it.
	defer c.disc.Close()

	iter := c.disc.RandomNodes()
	// Used to send the node from the iterator to this method.
	nodeCh := make(chan *enode.Node)
	// We need to run the iterator in another goroutine, because Next method
	// can be slow.
	go func() {
		defer close(nodeCh)
		for iter.Next() {
			nodeCh <- iter.Node()
		}
	}()

	for n := range nodeCh {
		// We have to directly request the ENR from the node to make sure that
		// the node is alive.
		nn, err := c.disc.RequestENR(n)
		if err != nil {
			// If it's not alive, log and skip to the next node. We don't have
			// to return an error here to the upper level in the call stack.
			log.Printf("found unalive node\tid=%s", n.ID().TerminalString())
			continue
		}
		log.Printf("found alive node\tid=%s", n.ID().TerminalString())
		// TODO: Send the node back to the caller
		_ = nn

	}

	iter.Close()
	return nil
}

// Run all the necessary steps to produce `c.disc`.
func (c *Crawler) setupDiscovery() error {
	cfg := discover.Config{
		PrivateKey: c.privateKey,
		Bootnodes:  c.bootstrapNodes,
	}
	// By putting the empty string, it will create a memory database instead
	// of a persistent database.
	db, err := enode.OpenDB("")
	if err != nil {
		return err
	}

	// Create a new local ethereum p2p node.
	ln := enode.NewLocalNode(db, cfg.PrivateKey)
	// Bind to some UDP port.
	addr := "0.0.0.0:0"
	socket, err := net.ListenPacket("udp4", addr)
	if err != nil {
		return err
	}
	usocket := socket.(*net.UDPConn)

	// SetFallbackIP and SetFallbackUDP set the last-resort IP address.
	// This address is used if no endpoint prediction can be made.
	uaddr := socket.LocalAddr().(*net.UDPAddr)
	if uaddr.IP.IsUnspecified() {
		ln.SetFallbackIP(net.IP{127, 0, 0, 1})
	} else {
		ln.SetFallbackIP(uaddr.IP)
	}
	ln.SetFallbackUDP(uaddr.Port)

	// ListenV5 listens on the given connection. It creates many goroutines to
	// handle events and incoming packets.
	c.disc, err = discover.ListenV5(usocket, ln, cfg)
	if err != nil {
		return err
	}
	return nil
}
