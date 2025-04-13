package main

import (
	"context"
	"fmt"

	pb "github.com/kanishkmehta29/grpc-basics/proto"
)

var cnt int = 1

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	fmt.Printf("Server responded %v time\n", cnt)
	cnt++
	return &pb.HelloResponse{
		Message: "hello",
	}, nil
}
