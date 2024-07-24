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

func handleAppNoCommand(data string, plain *bool) int {
	if data[0] == '!' {
		fn := strings.Split(data[1:], " ")

		if fn[0] == "plain" {
			*plain = !*plain
		}
		return -1
	}
	return 0
}

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

		if fn[0] == "genbufln" {
			ourint, err := strconv.Atoi(fn[2])
			if err != nil {
				panic(err)
			}
			return Genbuf(fn[1]+"\n", ourint)
		}
		if fn[0] == "genbuflrn" {
			ourint, err := strconv.Atoi(fn[2])
			if err != nil {
				panic(err)
			}
			return Genbuf(fn[1]+"\r\n", ourint)
		}
		if fn[0] == "genbuflrn2" {
			ourint, err := strconv.Atoi(fn[2])
			if err != nil {
				panic(err)
			}
			return Genbuf(fn[1]+"\r\n\r\n", ourint)
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
		if fn[0] == "r" {
			return "\r"
		}
		if fn[0] == "rn" {
			return "\r\n"
		}
		if fn[0] == "rn2" {
			return "\r\n\r\n"
		}
	}
	return data
}

func readFromserver(conn net.Conn, plain *bool) {
	for {
		buf := make([]byte, 1024)

		_, err := conn.Read(buf)
		if err != nil {
			log.Println("Error reading data from connection:", err)
			return
		}
		recvtext := string(buf)

		if *plain {
			fmt.Print(recvtext)
		} else {
			fmt.Print("[reply] ", utils.SeeLinebreak(recvtext))
		}
	}
}

func ReplEventloop(conn net.Conn) {

	var plain bool = false
	go readFromserver(conn, &plain)

	for {
		consoleReader := bufio.NewReaderSize(os.Stdin, 1)
		inputtext, _ := consoleReader.ReadString('\n')
		ascii := byte(inputtext[0])

		if ascii == 27 || ascii == 3 {
			fmt.Println("Exiting...")
			os.Exit(0)
		}

		ret := handleAppNoCommand(strings.TrimSuffix(inputtext, "\n"), &plain)

		if ret != -1 {
			commandresult := handleCommand(strings.TrimSuffix(inputtext, "\n"))

			conn.Write([]byte(commandresult))
		}

	}

}
