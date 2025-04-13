package main

import (
	"io"
	"log"

	pb "github.com/kanishkmehta29/grpc-basics/proto"
)

func (s *helloServer) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBidirectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("got request with name : %v", req.Name)
		res := &pb.HelloResponse{
			Message: "Hello " + req.Name,
		}
		err = stream.Send(res)
		if err != nil {
			return err
		}
	}
}
