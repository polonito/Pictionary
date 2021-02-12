package main

import (
	"flag"
	"fmt"
	"net"
)

func main() {

	fmt.Println("Hello ")

	var port string
	flag.StringVar(&port, "port", "1234", "usage")
	fmt.Println("Before parsing", port)
	flag.Parse()
	fmt.Println("After parsing", port)
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	//fmt.Println("kon")
	b := make([]byte, 256)

	n, err := conn.Read(b)_
	if err != nil { 
		panic(err)
	}
	var name string = string(b)
	fmt.Println(n, name[:n])
}
