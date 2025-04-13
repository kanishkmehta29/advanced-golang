package main

import (
	"log"

	pb "github.com/kanishkmehta29/grpc-basics/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main(){
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Fatalf("did not connect :%v",err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Names: []string{"Alice", "Bob", "Claire"},
	}

	// callSayHello(client)
	// callSayHelloServerStream(client,names)
	// callSayHelloClientStream(client,names)
	callSayHelloBiStream(client,names)

}