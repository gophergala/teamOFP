package main

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/crowdmob/goamz/sqs"
)

//NotificationMessage simple struct for sending data
type NotificationMessage struct {
	Event string `json:"event"`
	Value string `json:"values"`
}
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
	}

}

func processQueue(ch chan *sqs.Message) {
	for m := range ch {
		// log.Println("Processing Message: ", m)
		messagebody := map[string]interface{}{}
		err := json.Unmarshal([]byte(m.Body), &messagebody)
		if err != nil {
			log.Panic("the unmarshall plan")
		}
		if messagebody["command"] == "play_track" {
			if str, ok := messagebody["param"].(string); ok {
				setNextTrack(str)
			} else {
				log.Panic("was unable to set current track")
			}

		}
	}
}

func pollSystem(queue *sqs.Queue) {
	playerState := getPlayerState()
	track := getCurrentTrackID()
	timeLeft := int(getTimeLeft())
	getNextSong := true
	// log.Println("starting player state: ", playerState)
	for {
		time.Sleep(time.Second / 2)
		//check player state
		currentPlayerState := getPlayerState()
		currentTimeLeft := int(getTimeLeft())
		currentTrack := getCurrentTrackID()
		if playerState != currentPlayerState {
			message := NotificationMessage{"player_" + currentPlayerState, ""}
			err := pushMessage(queue, message)
			if err != nil {
				log.Println(err)
			}
			playerState = currentPlayerState
			log.Println("player state changed: ", currentPlayerState)
		}
		//check player duration - is track over
		if currentTimeLeft != timeLeft {
			// log.Println("New Time : ", currentTimeLeft)
			timeLeft = currentTimeLeft
			message := NotificationMessage{"time_left", strconv.Itoa(timeLeft)}
			pushMessage(queue, message)
			if timeLeft < 30 && getNextSong { //lock out period
				getNextSong = false
				message := NotificationMessage{"track_end", track}
				pushMessage(queue, message)
			}
		}

		if currentTrack != track {
			if !getNextSong {
				message := NotificationMessage{"track_start", nextTrack}
				pushMessage(queue, message)
			}
			getNextSong = true
			nextTrack := getNextTrack()
			setCurrentTrack("spotify:track:" + nextTrack)
			track = currentTrack
		}

	}
}

func pushMessage(q *sqs.Queue, message interface{}) error {
	// log.Println("message: ", message)
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
