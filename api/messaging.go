package main

import (
	"encoding/json"
	"github.com/crowdmob/goamz/sqs"
	"log"
)

type NotificationMessage struct {
	Event string `json:"event"`
	Value string `json:"values"`
}

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

		n := &NotificationMessage{}

		if err := json.Unmarshal([]byte(m.Body), &n); err != nil {
			log.Panic(err)
		}

		switch n.Event {
		case "track_end":
			log.Println("Song Ended")
			queueNextTrack()

		case "track_start":
			log.Println("Track Started: ", n.Value)

		case "player_paused":
			log.Println("Player Paused")

		case "player_playing":
			log.Println("Player Playing")

		case "player_stopped":
			log.Println("Player Stopped")

		case "time_left":
			log.Println("Time Left: ", n.Value)

		}
	}
}

func pushMessage(q *sqs.Queue, message interface{}) error {

	j, err := json.Marshal(message)
	if err != nil {
		return err
	}

	_, err = q.SendMessage(string(j))
	if err != nil {
		return err
	}

	return nil
}
