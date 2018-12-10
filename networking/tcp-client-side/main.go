package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// open connection
	conn, err := net.Dial("tcp", "localhost:50000")
	if err != nil {
		// no connection could be made as the target machine actively refused it
		fmt.Println("Error dialing", err.Error())
		return
	}

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please type your name:")
	clientName, _ := inputReader.ReadString('\n')
	trimmedClient := strings.Trim(clientName, "\n") // "\r\n" on windows, "\n" on linux

	// send info to server until Quit
	for {
		fmt.Println("What would you like to send to the server? Type Q to quit")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\n")
		if trimmedInput == "Q" {
			return
		}
		_, err = conn.Write([]byte(trimmedClient + " says: " + trimmedInput))
	}
}
