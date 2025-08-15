package main

import (
	pb "grpc-demo/proto"
	"io"
	"log"

	"google.golang.org/grpc"
)

func (s *helloServer) SayHelloClientStreaming(stream grpc.ClientStreamingServer[pb.HelloRequest, pb.MessageList]) error {
	var messages []string

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessageList{Messages: messages})
		}
		if err != nil {
			return err
		}
		log.Printf("got request with name: %v", req.Name)

		messages = append(messages, "Hello ", req.Name)
	}
}
