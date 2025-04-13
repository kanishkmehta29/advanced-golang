package main

import (
	"log"
	"net"
    pb "github.com/kanishkmehta29/grpc-basics/proto"
	"google.golang.org/grpc"
)

type helloServer struct{
	pb.GreetServiceServer
}

func main(){
	lis,err := net.Listen("tcp",":8080")
	if err !=nil{
		log.Fatalf("failed to start the server %v",err.Error())
	}

	//creates a new grpc server
	grpcServer := grpc.NewServer()

	//register with the greet service
	pb.RegisterGreetServiceServer(grpcServer,&helloServer{})
	log.Printf("server started at %v",lis.Addr())

	err = grpcServer.Serve(lis)
	if err != nil{
		log.Fatalf("failed to start %v",err.Error())
	}

}