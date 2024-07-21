package main

import (
	"fmt"
	"net"
)

func OpenTCPConnection(host string, port int) net.Conn {
	addr := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}

	return conn
}
