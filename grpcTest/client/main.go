package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "grpcTest/helloworld"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "sup world"
)

func main() {
	// 建立网络连接，创建grpc会话
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// 创建grpc客户端对象
	c := pb.NewSupGreeterClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// rpc调用
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("coudld not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
