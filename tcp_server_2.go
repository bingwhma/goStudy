package main

import (
	"fmt"
	"net"
	"os"
)

const (
	ip   = "127.0.0.1"
	port = 8888
)

func Server(listen *net.TCPListener) {
	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			fmt.Println("error :", err.Error())
			continue
		}
		fmt.Println("客户端连接来自:", conn.RemoteAddr().String())

		defer conn.Close()
		go func() {
			data := make([]byte, 128)
			for {
				in, err := conn.Read(data)
				if err != nil {
					fmt.Println("error :", err.Error())
					break
				}

				conn.Write([]byte{'O', 'K'})
			}
		}()
	}
}

func main() {
	listen, err := net.ListenTCP("tcp", &net.TCPAddr{net.ParseIP(ip), port, ""})
	if err != nil {
		os.Exit(1)
	}

	Server(listen)
}
