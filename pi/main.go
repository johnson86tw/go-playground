package main

import (
	"context"
	"log"
	"net"

	"github.com/chnejohnson/pi/proto"
	"google.golang.org/grpc"
)

type nodeService struct {
	proto.UnimplementedNodeServiceServer
}

func main() {
	var ns nodeService

	listener, err := net.Listen("tcp", ":4000")
	if err != nil {
		log.Fatalln("Fail to listen on port 4000", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterNodeServiceServer(grpcServer, ns)

	log.Println("server start up")
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("Fail to serve grpc server on port 4000: %+v\n", err)
	}

}

func (ns nodeService) UploadData(ctx context.Context, d *proto.SensorData) (*proto.Response, error) {
	log.Printf("Data: %+v\n", d)
	return &proto.Response{Msg: "success"}, nil
}
