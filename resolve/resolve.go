package main

import (
	"fmt"
	"net"
	"os"
)

type network string

const (
	IP   network = "ip"
	IPV4         = "ip4"
	IPV6         = "ip6"
)

func (n network) isValid() error {
	switch n {
	case IP, IPV4, IPV6:
		return nil
	}
	return fmt.Errorf("invalid IP type")
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s hostname ipAdressType\n", os.Args[0])
		os.Exit(1)
	}

	netType := network(os.Args[2])
	if netType.isValid() != nil {
		fmt.Fprintf(os.Stderr, "invalid ipAddress type")
	}

	name := os.Args[1]

	addr, err := net.ResolveIPAddr(string(netType), name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Resolution Error: %s", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Resolved address: %s\n", addr.String())
	os.Exit(0)
}
