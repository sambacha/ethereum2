package crawler

import (
	"errors"
	"fmt"
	"log"
	"testing"

	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ppopth/discv5-crawl/internal/testlog"
)

type nodeInfo struct {
	url   string
	alive bool
}

var (
	nodeInfos = []nodeInfo{
		{url: "enr:-KO4QDBsHwuYdxyb_KR_sJEt-5ikIsdfyQHK6zi72KiDXTIgDGf9mQl8hen6ycgbJyaSgjbe9_lLy6lcZZA5iwECoCWCATWEZXRoMpCvyqugAQAAAP__________gmlkgnY0gmlwhAMTwp2Jc2VjcDI1NmsxoQOGl6EENtmMz8v16Tr31ju-FQn54B0zJBb8WKXnbZjR84N0Y3CCIyiDdWRwgiMo", alive: true},
		{url: "enr:-Ku4QLylXZ0DWTelCTZQJxl2lsJFYYNk9B_Q2YXYfnxAiYCsRyOJnbVvxWRnQqiD1KTpa4YCdPwcdilx0ALtjIwLRjIHh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhDayLMaJc2VjcDI1NmsxoQK2sBOLGcUb4AwuYzFuAVCaNHA-dy24UuEKkeFNgCVCsIN1ZHCCIyg", alive: false},
		{url: "enr:-Ly4QKQ4BqHAOloSz-_lYVbfPpuAbn3uFxFiRSmWNzSEJZrsVnG-kTqjAleCu-KkSxvmIpt_ZIMmgUMbrWGdvDyEuM08h2F0dG5ldHOI__________-EZXRoMpDucelzYgAAcf__________gmlkgnY0gmlwhES3XM2Jc2VjcDI1NmsxoQK79EwWY2Zi9wvUKcFGkN3-VwoMvLLCJCKHQxFH6xgPyYhzeW5jbmV0cw-DdGNwgiMog3VkcIIjKA", alive: true},         // has attnets
		{url: "enr:-MK4QB-ycOj1GuzRW8pjXiMxRhQz0Yby-Z_KWwZ_D3ddGy2dbWOVTmj3E6hFkoFTGeey1qJhq2bddsSnMz9xvWNneKGGAX26RhZSh2F0dG5ldHOI-_-___7__9-EZXRoMpCvyqugAQAAAP__________gmlkgnY0gmlwhCPc-ZyJc2VjcDI1NmsxoQMJdU5g6WmwFY10zH2rB7qyM-3hBgPH9mTtRLu-zv1FIYhzeW5jbmV0cwCDdGNwgjLIg3VkcIIu4A", alive: true}, // has attnets
		{url: "enr:-IS4QDAyibHCzYZmIYZCjXwU9BqpotWmv2BsFlIq1V31BwDDMJPFEbox1ijT5c2Ou3kvieOKejxuaCqIcjxBjJ_3j_cBgmlkgnY0gmlwhAMaHiCJc2VjcDI1NmsxoQJIdpj_foZ02MXz4It8xKD7yUHTBx7lVFn3oeRP21KRV4N1ZHCCIyg", alive: false},
	}
)

func TestNew(t *testing.T) {
	c := &Config{}
	n := New(c)
	if n == nil {
		t.Error("New returns nil")
	}
	if len(n.bootstrapNodes) == 0 {
		t.Error("New doesn't create any bootstrap node")
	}
	if n.log == nil {
		t.Error("New doesn't create the logger")
	}
	if n.config != c {
		t.Error("New doesn't reference the config")
	}
}

func TestNewWithBootNodes(t *testing.T) {
	urls := make([]string, 0, len(nodeInfos))
	for _, info := range nodeInfos {
		urls = append(urls, info.url)
	}
	c := &Config{BootstrapUrls: urls}
	n := New(c)
	if n == nil {
		t.Error("New returns nil")
	}
	if len(n.bootstrapNodes) != len(urls) {
		t.Errorf("New creates the incorrect number of bootstrap nodes: got: %v expected %v", len(n.bootstrapNodes), len(urls))
	}
	if n.log == nil {
		t.Error("New doesn't create the logger")
	}
	if n.config != c {
		t.Error("New doesn't reference the config")
	}
}

type fakeDisc struct {
	// List of nodes.
	l []*enode.Node
	// List of alive node infos.
	a []nodeInfo
	// Used to look up if the node is alive.
	m map[*enode.Node]bool
}

func NewFakeDisc(infos []nodeInfo) *fakeDisc {
	urls := make([]string, 0, len(infos))
	for _, info := range infos {
		urls = append(urls, info.url)
	}
	nodes := parseNodeUrls(urls)
	m := make(map[*enode.Node]bool)
	for i, n := range nodes {
		m[n] = infos[i].alive
	}
	a := make([]nodeInfo, 0)
	for i := range nodes {
		if infos[i].alive {
			a = append(a, infos[i])
		}
	}
	return &fakeDisc{
		l: nodes,
		a: a,
		m: m,
	}
}
func (fd *fakeDisc) RandomNodes() enode.Iterator {
	return enode.IterNodes(fd.l)
}
func (fd *fakeDisc) RequestENR(n *enode.Node) (*enode.Node, error) {
	if fd.m[n] {
		return n, nil
	} else {
		return nil, errors.New("node is not alive")
	}
}
func (fd fakeDisc) Close() {}

func TestRun(t *testing.T) {
	disc := NewFakeDisc(nodeInfos)
	c := New(&Config{})
	logger := testlog.New()
	// Inject the discovery before Run.
	c.disc = disc
	c.log = log.New(logger, "", 0)
	ch := make(chan *enode.Node)
	go func() {
		defer close(ch)
		c.run(ch)
	}()

	// Check the returned nodes.
	// An index to iterate through all the alive nodes (disc.a).
	i := 0
	for n := range ch {
		if i >= len(disc.a) {
			t.Fatalf("the number of returned nodes exceeds expected: %v", len(disc.a))
		}
		nInfo := disc.a[i]
		i++
		if nInfo.url != n.String() {
			t.Errorf("got an incorrect node: got: %q expected: %q", n.String(), nInfo.url)
		}
	}
	// Check the log.
	if logger.Len() != len(disc.l) {
		t.Errorf("the number of log lines is incorrect: got: %v expected: %v", logger.Len(), len(disc.l))
	}
	for _, n := range disc.l {
		var expected string
		if disc.m[n] {
			expected = fmt.Sprintf("found alive node\t\tid=%s\n", n.ID().TerminalString())
		} else {
			expected = fmt.Sprintf("found unalive node\t\tid=%s\n", n.ID().TerminalString())
		}
		if !logger.Has(expected) {
			t.Errorf("not found expected log: %q", expected)
		}
	}
}
