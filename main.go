package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"strings"
)

func server() {
	//listen on port
	ln, err := net.Listen("tcp", ":9001")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleServerConnection(c)
	}
}

func handleServerConnection(c net.Conn) {
	msg := ""
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("received", msg)
	}

	defer c.Close()

}

func client(msg string) {
	// connect to server
	client, err := net.Dial("tcp", "localhost:9001")
	if err != nil {
		fmt.Println(err)
		return
	}

	// send message
	fmt.Println("sending..")
	err = gob.NewEncoder(client).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}
	defer client.Close()
}

func main() {
	go server()
	for {
		var input string
		fmt.Scanln(&input)
		if strings.Compare(input, "1") == 0 {
			return
		}
		client(input)

	}
}
