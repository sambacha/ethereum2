package crawler

import (
	"log"

	"github.com/ethereum/go-ethereum/p2p/enode"
)

func parseBoostrapUrls(urls []string) []*enode.Node {
	nodes := make([]*enode.Node, 0, len(urls))
	for _, url := range urls {
		if url != "" {
			node, err := enode.Parse(enode.ValidSchemes, url)
			if err != nil {
				log.Printf("bootstrap URL invalid enode %s err %s", url, err)
				continue
			}
			nodes = append(nodes, node)
		}
	}
	return nodes
}
