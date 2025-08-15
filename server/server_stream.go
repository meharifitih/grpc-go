package main

import (
	pb "grpc-demo/proto"
	"log"
	"time"

	"google.golang.org/grpc"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NameList, stream grpc.ServerStreamingServer[pb.HelloResponse]) error {
	log.Printf("got request wit names : %v", req.Names)

	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello " + name,
		}

		if err := stream.Send(res); err != nil {
			return err
		}

		time.Sleep(time.Second * 2)
	}

	return nil
}
