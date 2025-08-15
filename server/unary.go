package main

import (
	"context"
	pb "grpc-demo/proto"
)

func (s *helloServer) SayHello(context.Context, *pb.NoParam) (*pb.HelloResponse, error){
	return  &pb.HelloResponse{
		Message: "Hello",
	},nil
}
