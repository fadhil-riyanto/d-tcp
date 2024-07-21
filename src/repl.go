package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func ReplEventloop(conn net.Conn) {
	for {
		consoleReader := bufio.NewReaderSize(os.Stdin, 1)
		inputtext, _ := consoleReader.ReadString('\n')
		ascii := byte(inputtext[0])

		if ascii == 27 || ascii == 3 {
			fmt.Println("Exiting...")
			os.Exit(0)
		}

		// fmt.Println("$ ", inputtext)
		conn.Write([]byte(strings.TrimSuffix(inputtext, "\n")))

	}

}
