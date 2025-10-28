package main

import (
	"RPCTest/server"
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	log.Println("已连接到服务端")
	args := server.Args{A: 99, B: 1}
	var reply int
	err = client.Call("MathService.Add", args, &reply)
	if err != nil {
		log.Fatal("add:", err)
	}
	fmt.Printf("Add %d + %d = %d\n", args.A, args.B, reply)
}
