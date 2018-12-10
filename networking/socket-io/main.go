package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	var (
		host   = "www.apache.org"
		port   = "80"
		remote = host + ":" + port
		msg    = "GET / \n"
		data   = make([]uint8, 4096)
		read   = true
		count  = 0
	)

	// create the socket
	con, err := net.Dial("tcp", remote)

	// send an HTTP GET request
	io.WriteString(con, msg)

	// read the response from the web server
	for read {
		count, err = con.Read(data)
		read = (err == nil)
		fmt.Printf(string(data[0:count]))
	}
	con.Close()
}
