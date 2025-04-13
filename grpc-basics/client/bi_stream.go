package main

import (
	"context"
	"io"
	"log"
	"time"
	pb "github.com/kanishkmehta29/grpc-basics/proto"
)

func callSayHelloBiStream(client pb.GreetServiceClient,names *pb.NamesList){
	log.Println("Started streaming from client side")
	stream,err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil{
		log.Fatalf("error in connecting stream : %v",err)
	}
    waitc := make(chan struct{})

	go func(){
		for{
		message,err := stream.Recv()
		if err == io.EOF{
			break
		}
		if err != nil{
			log.Fatalf("error while accepting stream : %v",err)
		}
		log.Println(message)
	}
	close(waitc)
	}()

	for _,name := range names.Names{
		req := &pb.HelloRequest{
			Name:name,
		}
		err = stream.Send(req)
		if err != nil{
			log.Fatalf("error while sending : %v",err)
		}
		time.Sleep(2*time.Second)
	}
	stream.CloseSend()
	<-waitc
	log.Printf("Bidirectional streaming finished")
}