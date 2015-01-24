package main

import (
	"log"
	"strconv"
)

func getTimeLeft() int {
	duration, derr := strconv.Atoi(systemCall("duration", ""))
	position, perr := strconv.Atoi(systemCall("position", ""))

	if derr != nil {
		log.Panic(derr)
	}

	if perr != nil {
		log.Panic(perr)
	}

	return duration - position
}
