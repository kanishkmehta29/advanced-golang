package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"github.com/gofiber/fiber/v2"
)

type Comment struct {
	Text string `form:"text" json:"text"`
}

func main() {
	app := fiber.New()
	api := app.Group("/api/v1")
	api.Post("/comments", createComment)

	app.Listen(":8080")
}

func createComment(ctx *fiber.Ctx) error {

	var cmt Comment

	err := ctx.BodyParser(&cmt)
	if err != nil {
		log.Printf("error in parsing body of comment : %v", err)
		ctx.Status(400).JSON(fiber.Map{
			"sucess":  false,
			"message": err.Error(),
		})
		return err
	}

	cmtInBytes, err := json.Marshal(cmt)
	if err != nil {
		log.Printf("error in marshalling comment : %v", err)
		ctx.Status(400).JSON(fiber.Map{
			"sucess":  false,
			"message": err.Error(),
		})
		return err
	}

	err = PushCommentToQueue("comments", cmtInBytes)
	if err != nil {
		log.Printf("error pushing comment to queue : %v", err)
		ctx.Status(400).JSON(fiber.Map{
			"sucess":  false,
			"message": err.Error(),
		})
		return err
	}
	return err
}

func PushCommentToQueue(topic string, message []byte) error {
	brokersUrl := []string{"localhost:29092"}
	producer, err := ConnectProducer(brokersUrl)
	if err != nil {
		return err
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return nil
	}

	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)

	return nil

}

func ConnectProducer(brokersUrl []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	conn, err := sarama.NewSyncProducer(brokersUrl, config)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
