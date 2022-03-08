package main

import (
	"bytes"
	"flag"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/p2p/enr"
	"github.com/ppopth/discv5-crawl/attnets"
	"github.com/ppopth/discv5-crawl/crawler"
)

var (
	bootnodesFlag = flag.String("bootnodes", "", "Comma separated nodes used for bootstrapping")
)

func main() {
	flag.Parse()
	log.Println("started discv5-crawl")

	var bootUrls []string
	if *bootnodesFlag != "" {
		bootUrls = strings.Split(*bootnodesFlag, ",")
	}
	nodeCh := make(chan *enode.Node)
	// Run the crawler in another goroutine and send the discovered nodes
	// through a channel.
	go func() {
		defer close(nodeCh)
		// If bootUrls is nil, the crawler will use the default boot nodes.
		cr := crawler.New(bootUrls)
		if err := cr.Run(nodeCh); err != nil {
			log.Fatal(err)
		}
	}()

	for n := range nodeCh {
		var e attnets.Attnets
		if err := n.Load(&e); err != nil {
			// If there is an error, it probably means the key "attnets" is not
			// present. In which case, we silently skip. Otherwise, we log the
			// error.
			if !enr.IsNotFound(err) {
				log.Printf("found bad node: %s", err)
			}
			continue
		}
		if bytes.Equal(e[:], make([]byte, 8)) {
			// If attnets is composed of all zeros, we don't consider it.
			continue
		}
		log.Printf("found node with attnets\tid=%s\tattnets=%s", n.ID().TerminalString(), e)
	}
}
