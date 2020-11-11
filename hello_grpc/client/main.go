package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/hello-grpc/proto"
	"google.golang.org/grpc"
)

const (
	defaultName = "john"
)

type client struct{}

func main() {
	// what's grpc.WithBlock() do?
	conn, err := grpc.Dial("localhost:4000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot connect to port 4000 %v\n", err)
	}
	defer conn.Close()

	c := proto.NewGreeterServiceClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	// what exactly is context do?
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.Greet(ctx, &proto.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("fail to greet: %v\n", err)
	}

	log.Printf("Greeting from server: %s\n", res.GetReply())

}
