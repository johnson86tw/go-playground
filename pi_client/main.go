package main

import (
	"context"
	"log"

	"github.com/chnejohnson/pi_client/proto"
	"google.golang.org/grpc"
)

var temp float32 = 23.5
var hmd float32 = 87.6

func main() {
	conn, err := grpc.Dial("localhost:4000", grpc.WithInsecure())
	if err != nil {
		log.Panicf("cannot connect to port 4000 %+v\n", err)
	}

	defer conn.Close()

	cli := proto.NewNodeServiceClient(conn)

	d := &proto.SensorData{
		Temp: temp,
		Hmd:  hmd,
	}

	res, err := cli.UploadData(context.Background(), d)
	if err != nil {
		log.Fatalf("Fail to upload: %+v\n", err)
	}

	log.Println("Server reply:", res.GetMsg())
}
