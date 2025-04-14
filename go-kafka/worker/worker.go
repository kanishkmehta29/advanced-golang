package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

func main(){
	topic := "comments"
	worker, err := connectConsumer([]string{"localhost:29092"})
	if err != nil{
		log.Fatalf("error in connecting to consumer :%v",err)
	}

	consumer,err := worker.ConsumePartition(topic,0,sarama.OffsetOldest)
	if err != nil{
		log.Fatalf("error in consumer partition :%v",err)
	}

	fmt.Println("consumer started")
	sigchan := make(chan os.Signal,1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	msgCount := 0

	doneCh := make(chan struct{})
	go func(){
		for{
			select{
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				msgCount++
				fmt.Printf("Received message Count %d: | Topic(%s) | Message(%s) \n", msgCount, string(msg.Topic), string(msg.Value))
			case <-sigchan:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}	
			}
		}
	}()

	<-doneCh
	fmt.Println("Processed",msgCount,"messages")
	err = worker.Close()
	if err != nil{
		log.Fatal(err)
	}
}

func connectConsumer(brokersUrl []string) (sarama.Consumer,error){
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	conn,err := sarama.NewConsumer(brokersUrl,config)
	if err != nil{
		return nil,err
	}
	return conn,nil
}