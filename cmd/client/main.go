package main

import (
	"flag"
	"fmt"
	"net"
)

func main() {
	var host string
	var port string
	var name string
	flag.StringVar(&host, "host", "127.0.0.1", "host to connect to")
	flag.StringVar(&port, "port", "1234", "port to connect to")
	flag.StringVar(&name, "name", "", "name to be greeted")

	flag.Parse()

	fmt.Println(host, port, name)
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		panic(err)
	}
	fmt.Println("connected")
	n, err := conn.Write([]byte(name))
	if err != nil {
		panic(err)
	}
	fmt.Println(n, name[:n])
}
