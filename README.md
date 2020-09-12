[![CI](https://github.com/takumin/gopacket-radius/workflows/CI/badge.svg)](https://github.com/takumin/gopacket-radius/actions)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/takumin/gopacket-radius)](https://pkg.go.dev/github.com/takumin/gopacket-radius)

# gopacket-radius
[google/gopacket: Provides packet processing capabilities for Go](https://github.com/google/gopacket) extension

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
	_ "github.com/takumin/gopacket-radius/pkg/radius"
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
