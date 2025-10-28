package main

import (
	"RPCTest/server"
	"log"
	"net"
	"net/rpc"
)

func main() {
	rpc.RegisterName("MathService", new(server.MathServer))
	log.Println("开始监听")
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	} else {
		rpc.Accept(l)
	}
}
