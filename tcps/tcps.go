package main

import "net"

// call over wifi
// * export ip=ifconfig | grep "inet " | grep -v 127.0.0.1
// * nc $ip 1201
func main() {
	service := ":1201"

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		panic(err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		// don't block handle per connection on its own routine
		go handle(conn)
	}

}

func handle(conn net.Conn) {
	defer conn.Close() // close on exit

	// var buf [512]byte

	// for {
	// read from connection
	// n, err := conn.Read(buf[0:])
	// if err != nil {
	// 	return
	// }

	// write back the bytes
	_, err := conn.Write([]byte("piggs !"))
	if err != nil {
		return
	}
	// }
}
