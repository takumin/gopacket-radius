[![CI](https://github.com/takumin/gopacket-radius/workflows/CI/badge.svg)](https://github.com/takumin/gopacket-radius/actions)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/takumin/gopacket-radius)](https://pkg.go.dev/github.com/takumin/gopacket-radius)
[![codecov](https://codecov.io/gh/takumin/gopacket-radius/branch/master/graph/badge.svg)](https://codecov.io/gh/takumin/gopacket-radius)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Ftakumin%2Fgopacket-radius.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Ftakumin%2Fgopacket-radius?ref=badge_shield)

# gopacket-radius
[RADIUS Protocol](https://en.wikipedia.org/wiki/RADIUS) extension for [google/gopacket](https://github.com/google/gopacket)

# How to use
```go
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	_ "github.com/takumin/gopacket-radius"
)

func main() {
	var (
		f string
	)

	flag.StringVar(&f, "pcap", "", "pcap file path")
	flag.Parse()

	if _, err := os.Stat(f); err != nil {
		log.Fatal(err)
	}

	h, err := pcap.OpenOffline(f)
	if err != nil {
		log.Fatal(err)
	}
	defer h.Close()

	src := gopacket.NewPacketSource(h, h.LinkType())
	for {
		pkt, err := src.NextPacket()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		fmt.Println(pkt)
	}
}
```


## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Ftakumin%2Fgopacket-radius.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Ftakumin%2Fgopacket-radius?ref=badge_large)