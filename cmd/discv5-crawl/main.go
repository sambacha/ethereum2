package main

import (
	"log"

	"github.com/ppopth/discv5-crawl/crawler"
)

func main() {
	log.Println("started discv5-crawl")
	cr := crawler.New()
	if err := cr.Run(); err != nil {
		log.Fatal(err)
	}
}
