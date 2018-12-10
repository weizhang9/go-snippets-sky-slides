package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	// net.Dial (👌🏼) abstract the network family and transport, so ipv4/6, tcp and udp all share a common interface Conn
	// a Conn interface type is return when net.Dial into a remote system using
	// Conn interface type will be used to send and receive data

	conn, err := net.Dial("tcp", "192.0.32.10:80") // tcp ipv4
	checkConnection(conn, err)

	conn, err = net.Dial("udp", "192.0.32.10:80")
	checkConnection(conn, err)

	conn, err = net.Dial("tcp", "[2620:0:2d0:200::10]:80") // tcp ipv6
	checkConnection(conn, err)
}

func checkConnection(conn net.Conn, err error) {
	if err != nil {
		fmt.Printf("error %v connecting", conn)
		os.Exit(1)
	}
	fmt.Printf("Connection is made with %v", conn)
}
