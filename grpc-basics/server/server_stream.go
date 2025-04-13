package main

import (
	"log"
	"time"
	pb "github.com/kanishkmehta29/grpc-basics/proto"
)

func (s *helloServer)SayHelloServerStreaming(req *pb.NamesList,stream pb.GreetService_SayHelloServerStreamingServer) error{
	log.Printf("Got request with names : %v",req.Names)
	for _,name := range req.Names{
		res := &pb.HelloResponse{
			Message: "Hello " + name,
		}
		err := stream.Send(res)
		if err != nil{
            return err
		}
		time.Sleep(2*time.Second)
	}
	return nil
}