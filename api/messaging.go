package main

import (
	"github.com/crowdmob/goamz/sqs"
	"log"
)

func listenOnQueue(q *sqs.Queue, ch chan *sqs.Message) {

	for {
		resp, err := q.ReceiveMessage(1)
		if err != nil {
			log.Panic(err)
		}

		for _, m := range resp.Messages {
			ch <- &m
			q.DeleteMessage(&m)
		}
	}

}

func processQueue(ch chan *sqs.Message) {
	for m := range ch {
		log.Println("Processing Message: ", m)
	}
}

//func pushMessage()
