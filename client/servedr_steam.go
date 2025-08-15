package main

import (
	"context"
	pb "grpc-demo/proto"
	"io"
	"log"
	"time"
)

func callSayHelloServerStreaming(client pb.GreetServiceClient, names *pb.NameList) {
	log.Println("streaming started")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()

	stream, err := client.SayHelloServerStreaming(ctx, names)
	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while streaming: %v", err)
		}

		log.Println(message)
	}
	log.Println("streaming finished")
}
