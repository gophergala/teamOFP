package main

import (
	"log"
	"strconv"
)

var nextTrack string

func getTimeLeft() float32 {
	duration, derr := strconv.ParseFloat(systemCall("duration", ""), 64)
	position, perr := strconv.ParseFloat(systemCall("position", ""), 64)

	if derr != nil {
		log.Panic(derr)
	}

	if perr != nil {
		log.Panic(perr)
	}

	return float32(duration - position)
}

func getPlayerState() string {
	return systemCall("state", "")
}

func getCurrentTrack() string {
	return systemCall("name", "")
}

func getCurrentTrackID() string {
	return systemCall("id", "")
}

func setCurrentTrack(id string) {
	systemCall("play_track", "\""+id+"\"")
}

func getNextTrack() string {
	return nextTrack
}

func setNextTrack(track string) {
	nextTrack = track
}
