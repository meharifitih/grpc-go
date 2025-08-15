package main

import (
	"context"
	pb "grpc-demo/proto"
	"io"
	"log"
	"time"
)

func saySatHelloBiDirectional(client pb.GreetServiceClient, names pb.NameList) {
	log.Println("Bidirectional streaming has started")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	stream, err := client.SatHelloBiDirectional(ctx)
	if err != nil {
		log.Fatalf("could not send names: %v", names.Names)
	}

	waitc := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal("error while streaming: %v", err)
			}
			log.Println(message)
		}

		close(waitc)
	}()

	for _, name := range names.Names {
		pb := &pb.HelloRequest{
			Name: name,
		}

		if err := stream.Send(pb); err != nil {
			log.Fatalf("error while sending %v", err)
		}

		time.Sleep(2 * time.Second)
	}

	stream.CloseSend()
	<-waitc
	log.Printf("Bidirectional stream finished")
}
