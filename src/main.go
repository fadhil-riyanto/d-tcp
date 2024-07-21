package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type prop struct {
	host string
	port int16
}

func showhelp() {
	fmt.Printf("use ./dtcp <ADDRESS> <HOST>")
}

func getProp(prop *prop) error {
	// fmt.Println(len(os.Args))
	if len(os.Args) != 3 {
		showhelp()
		return errors.New("invalid length")
	}
	prop.host = os.Args[1]
	res, err := strconv.Atoi(os.Args[2])
	if err != nil {
		showhelp()
		return errors.New("invalid type")
	}

	prop.port = int16(res)
	return nil
}

func main() {
	var prop prop
	err := getProp(&prop)
	if err != nil {
		os.Exit(0)
	}

	fmt.Println(prop.host)
	fmt.Println(prop.port)
}
