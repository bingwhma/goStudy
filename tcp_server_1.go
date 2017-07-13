package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	tcp_addr, err := net.ResolveTCPAddr("tcp4", ":8888")
	if err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}

	tcp_listener, err := net.ListenTCP("tcp", tcp_addr)
	if err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}

	str_header := "Hello, thi is xxx"
	for {
		//waiting connection from client
		tcp_connection, err := tcp_listener.Accept()
		if err != nil {
			continue
		}

		//get the current time infor
		str_time := time.Now().String()
		//set greeting words for client connection
		str_greetings := str_header + str_time

		//send information to client
		tcp_connection.Write([]byte(str_greetings))
		tcp_connection.Close()
	}
}
