package main

import (
	"context"
	"log"
	"net"

	pb "../helloworld"
	"google.golang.org.grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedSupGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, err) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Sup say Hello to " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("fail to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSupGreeterServer(s, &server{})
	if err := s.Server(lis); err != nil {
		log.Fatalf("fail to serve: %v",err)
	}
}