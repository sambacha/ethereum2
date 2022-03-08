package main

import (
	"flag"
	"log"
	"strings"

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
	// If bootnodes is nil, the crawler will use the default bootstrap nodes.
	cr := crawler.New(bootUrls)
	if err := cr.Run(); err != nil {
		log.Fatal(err)
	}
}
