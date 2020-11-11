package main

import (
	"context"
	"log"
	"net"

	"github.com/hello-grpc/proto"
	"google.golang.org/grpc"
)

type service struct {
}

func main() {
	var s service

	listener, err := net.Listen("tcp", ":4000")
	if err != nil {
		log.Fatalf("Fail to listen on port 4000: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterGreeterServiceServer(grpcServer, s)

	log.Println("server start on port 4000")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Fail to serve grpc server on port 4000: %v", err)
	}

}

// Greet implements GreeterServiceServer interface
func (s service) Greet(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	log.Printf("Recevice %s\n", req.GetName())
	return &proto.HelloResponse{Reply: "Hello, " + req.GetName()}, nil
}
