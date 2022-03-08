package main

import (
	"flag"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/p2p/enode"
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
	var bootNodes []*enode.Node
	for _, url := range bootUrls {
		bootNodes = append(bootNodes, enode.MustParse(url))
	}
	nodeCh := make(chan *enode.Node)
	// Run the crawler in another goroutine and send the discovered nodes
	// through a channel.
	go func() {
		defer close(nodeCh)
		// If bootUrls is nil, the crawler will use the default boot nodes.
		cfg := &crawler.Config{BootNodes: bootNodes}
		cr := crawler.New(cfg)
		if err := cr.Run(nodeCh); err != nil {
			log.Fatal(err)
		}
	}()

	for n := range nodeCh {
		// Inspect the node and probably emit some log according to its
		// attributes.
		inspect(n)
	}
}
