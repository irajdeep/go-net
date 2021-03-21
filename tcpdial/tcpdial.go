package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		panic(err.Error())
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		panic(err.Error())
	}

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	if err != nil {
		panic(err.Error())
	}

	result, err := ioutil.ReadAll(conn)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(string(result))

}
