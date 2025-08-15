package main

import (
	"context"
	pb "grpc-demo/proto"
	"log"
	"time"
)

func saySayHelloClientStreaming(client pb.GreetServiceClient, names *pb.NameList) {
	log.Println("client streaming started")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	stream, err := client.SayHelloClientStreaming(ctx)
	if err != nil {
		log.Fatalf("could not send the names: %v", err)
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}

		if err := stream.Send(req); err != nil {
			log.Fatalf("error while sending %v", err)
		}
		log.Printf("sent the request with name %s", name)
		time.Sleep(time.Second * 2)
	}

	res, err := stream.CloseAndRecv()
	log.Println("client stream finished")
	if err != nil {
		log.Fatalf("error while receiving %v", err)
	}

	log.Printf("%v", res.Messages)
}
