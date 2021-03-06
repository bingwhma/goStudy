package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Server struct {
	ServerName string
	ServerIP   string
}
type Serverslice struct {
	Servers   []Server
	ServersID string
}

func main() {
	var s Serverslice
	var newServer Server
	newServer.ServerName = "Guangzhou_VPN"
	newServer.ServerIP = "127.0.0.1"
	s.Servers = append(s.Servers, newServer)
	s.Servers = append(s.Servers, Server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.2"})
	s.Servers = append(s.Servers, Server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.3"})
	s.ServersID = "team1"

	fmt.Println("Serverslice: \n", s)

	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}

	fmt.Println("json: \n", b)

	body := bytes.NewBuffer([]byte(b))

	fmt.Println("bytes.NewBuffer: \n", body)

	res, err := http.Post("http://localhost:9001/xiaoyue", "application/json;charset=utf-8", body)
	if err != nil {
		log.Fatal(err)
		return
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%s", result)
}
