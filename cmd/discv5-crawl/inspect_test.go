package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ppopth/discv5-crawl/internal/testlog"
)

type nodeInfo struct {
	url     string
	attnets string
}

var (
	nodeInfos = []nodeInfo{
		{url: "enr:-MK4QIRHEJ4TPuUKsOCYxsF6SYyKCOWKmLJVkxRGWfDdCKQpNcZGgSDneII0SJg7gFYUSJX0GiSQQK_KFecU0P2XjfaGAX8IHq4Rh2F0dG5ldHOIgAIAAAAAEACEZXRoMpCvyqugAgAAAP__________gmlkgnY0gmlwhDbbfwuJc2VjcDI1NmsxoQJizq1tU46iLN4eQdYOANP6WiuV4EopWaSKyTVgM_0vi4hzeW5jbmV0cwCDdGNwgjLIg3VkcIIu4A", attnets: "8002000000001000"},
		{url: "enr:-MK4QGny7WoJWxQx5POEw_4myGX2sCN3ga5W1Q_440tcr4vEEe0-gUgVgnpyBtMKkbBwJEYfmKTa9xpTOcNGnmDwqJeGAX7EcZoBh2F0dG5ldHOIgAwAAIACEAqEZXRoMpCC9KcrAgAQIP__________gmlkgnY0gmlwhCPgjeOJc2VjcDI1NmsxoQPp5_MyWky9d93GLTNk7paPOqkI-MrUYV8X52D2GOlGOIhzeW5jbmV0cwCDdGNwgjLIg3VkcIIu4A", attnets: "800c00008002100a"},
		{url: "enr:-Ku4QLylXZ0DWTelCTZQJxl2lsJFYYNk9B_Q2YXYfnxAiYCsRyOJnbVvxWRnQqiD1KTpa4YCdPwcdilx0ALtjIwLRjIHh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhDayLMaJc2VjcDI1NmsxoQK2sBOLGcUb4AwuYzFuAVCaNHA-dy24UuEKkeFNgCVCsIN1ZHCCIyg"},
		{url: "enr:-Iu4QCBao0bdWJeuJ2dfKh3D3QETA5_DO1w2Dg164X1bIDTye4mPS-ovRXJSQLNloMBCkZ7vnDuUv6v0MKTlKf6gYL4vgmlkgnY0gmlwhIeUK_uJc2VjcDI1NmsxoQM6CGkRvH4Epsdr6gV2S5HhvVzx5CoY3jU50munS8j-3oN0Y3CCdmeDdWRwgnZn"},
	}
)

func TestInspect(t *testing.T) {
	for _, info := range nodeInfos {
		node := enode.MustParse(info.url)
		logger := testlog.New()
		lg = log.New(logger, "", 0)
		inspect(node)
		expectedLogLen := 0
		if info.attnets != "" {
			expected := fmt.Sprintf("found node with attnets\tid=%s\tattnets=%s\n", node.ID().TerminalString(), info.attnets)
			if !logger.Has(expected) {
				t.Errorf("not found expected log: %q", expected)
			}
			expectedLogLen = 1
		}
		if logger.Len() != expectedLogLen {
			t.Errorf("the number of log lines is incorrect: got: %v expected: %v", logger.Len(), expectedLogLen)
		}
	}
}
