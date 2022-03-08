package main

import (
	"bytes"
	"log"

	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/p2p/enr"
	"github.com/ppopth/discv5-crawl/attnets"
)

// Declared to inject the custom log in the unit tests.
var lg = log.Default()

func inspect(n *enode.Node) {
	var e attnets.Attnets
	if err := n.Load(&e); err != nil {
		// If there is an error, it probably means the key "attnets" is not
		// present. In which case, we silently skip. Otherwise, we log the
		// error.
		if !enr.IsNotFound(err) {
			lg.Printf("found bad node: %s", err)
		}
		return
	}
	if bytes.Equal(e[:], make([]byte, 8)) {
		// If attnets is composed of all zeros, we don't consider it.
		return
	}
	lg.Printf("found node with attnets\tid=%s\tattnets=%s", n.ID().TerminalString(), e)
}
