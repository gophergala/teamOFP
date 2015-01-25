package main

import (
	"log"
	"strconv"
	"time"

	"github.com/crowdmob/goamz/sqs"
)

// polling sleep time
const sleepTime = time.Second / 2

func polling(queue *sqs.Queue) {
	playerState := getPlayerState()
	track := getCurrentTrackID()
	timeLeft := int(getTimeLeft())
	getNextSong := true
	//inner for loop variables
	var (
		currentPlayerState string
		currentTimeLeft    int
		currentTrack       string
	)
	// log.Println("starting player state: ", playerState)
	for {
		time.Sleep(sleepTime)
		//check player state
		currentPlayerState = getPlayerState()
		currentTimeLeft = int(getTimeLeft())
		currentTrack = getCurrentTrackID()
		// log.Println(currentTrack)
		if playerState != currentPlayerState {
			message := NotificationMessage{"player_" + currentPlayerState, "", currentTrack}
			err := pushMessage(queue, message)
			if err != nil {
				log.Println(err)
			}
			playerState = currentPlayerState
			// log.Println("player state changed: ", currentPlayerState)
		}
		//check player duration - is track over
		if currentTimeLeft != timeLeft {
			// log.Println("New Time : ", currentTimeLeft)
			timeLeft = currentTimeLeft
			message := NotificationMessage{"time_left", strconv.Itoa(timeLeft), currentTrack}
			pushMessage(queue, message)
			if timeLeft < 30 && getNextSong { //lock out period
				getNextSong = false
				message := NotificationMessage{"track_end", track, currentTrack}
				pushMessage(queue, message)
			}
		}

		if currentTrack != track {
			if !getNextSong {
				message := NotificationMessage{"track_start", nextTrack, nextTrack}
				pushMessage(queue, message)
			}
			getNextSong = true
			nextTrack := getNextTrack()
			setCurrentTrack(nextTrack)
			track = currentTrack
		}

	}
}
