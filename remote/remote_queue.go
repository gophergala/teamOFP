package main

import (
	"log"
	"time"

	"github.com/crowdmob/goamz/sqs"
)

// 	log.Println("Starting sqs processor")
//
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
//
// 	c.AWSAccess = os.Getenv("aws_access")
// 	c.AWSSecret = os.Getenv("aws_secret")
//
// 	done := make(chan bool)
// 	messageQueue := make(chan *sqs.Message)
//
// 	go listenOnQueue("dev", messageQueue)
// 	go processQueue(messageQueue)
//
// 	<-done

type context struct {
	AWSAccess string
	AWSSecret string
}

var c = &context{}

func listenOnQueue(queue string, ch chan *sqs.Message) {

	// Setup Queue
	s, err := sqs.NewFrom(c.AWSAccess, c.AWSSecret, "us-east-1")
	if err != nil {
		log.Panic(err)
	}
	q, err := s.GetQueue(queue)
	if err != nil {
		log.Panic(err)
	}

	for {
		resp, err := q.ReceiveMessage(1)
		if err != nil {
			log.Panic(err)
		}

		for _, m := range resp.Messages {
			ch <- &m
			q.DeleteMessage(&m)
		}
		time.Sleep(time.Second)
	}

}

func processQueue(ch chan *sqs.Message) {
	for m := range ch {
		log.Println("Processing Message: ", m)
	}
}

// func(queue)
