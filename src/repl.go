package main

import (
	"bufio"
	"dtcp/utils"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func handleCommand(data string) string {
	if data[0] == '/' {
		fn := strings.Split(data[1:], " ")

		if fn[0] == "genbuf" {
			ourint, err := strconv.Atoi(fn[2])
			if err != nil {
				panic(err)
			}
			return Genbuf(fn[1], ourint)
		}

		if fn[0] == "genrandbuf" {
			ourint, err := strconv.Atoi(fn[1])
			if err != nil {
				panic(err)
			}
			return Genrandbuf(ourint)
		}

		// reserved
		if fn[0] == "n" {
			return "\n"
		}
		if fn[0] == "nr" {
			return "\n\r"
		}
		if fn[0] == "nr2" {
			return "\n\r\r\n"
		}
	}
	return data
}

func readFromserver(conn net.Conn) {
	for {
		buf := make([]byte, 1024)

		_, err := conn.Read(buf)
		if err != nil {
			log.Println("Error reading data from connection:", err)
			return
		}
		recvtext := string(buf)
		fmt.Print("[reply] ", utils.SeeLinebreak(recvtext))
	}
}

func ReplEventloop(conn net.Conn) {
	go readFromserver(conn)

	for {
		consoleReader := bufio.NewReaderSize(os.Stdin, 1)
		inputtext, _ := consoleReader.ReadString('\n')
		ascii := byte(inputtext[0])

		if ascii == 27 || ascii == 3 {
			fmt.Println("Exiting...")
			os.Exit(0)
		}

		commandresult := handleCommand(strings.TrimSuffix(inputtext, "\n"))

		conn.Write([]byte(commandresult))

	}

}
