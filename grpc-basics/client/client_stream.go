package main

import (
	"context"
	"log"
	"time"
	pb "github.com/kanishkmehta29/grpc-basics/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient,names *pb.NamesList){
	log.Println("started streaming names from client side")
	stream,err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}

	for _, name := range names.Names{
		req := &pb.HelloRequest{
			Name:name,
		}
		err = stream.Send(req)
		if err != nil{
			log.Fatalf("error while sending names : %v",err.Error())
		}
		log.Printf("Send request with name: %v",name)
		time.Sleep(2*time.Second)
	}

	res,err := stream.CloseAndRecv()
	log.Println("Client streaming finished")
	if err != nil{
		log.Fatalf("Error while receiving %v",err.Error())
	}
	log.Printf("%v\n",res.Messages)
}