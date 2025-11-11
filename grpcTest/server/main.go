package main

import (
	"context"
	"log"
	"net"

	pb "grpcTest/helloworld"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// 封装rpc服务对象
type server struct {
	pb.UnimplementedSupGreeterServer
}

// 封装函数
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Sup say Hello to " + in.GetName()}, nil
}

func main() {
	// 创建一个监听端口
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("fail to listen: %v", err)
	}
	s := grpc.NewServer()
	// 注册rpc服务
	pb.RegisterSupGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("fail to serve: %v", err)
	}
}
