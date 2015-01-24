package main

import (
	"log"
	"time"
)

type TrackQueue struct {
	tracks    []string
	updatedAt time.Time
}

func (t *TrackQueue) push(track string) (int, error) {
	log.Println("Track pushed to track queue: ", track)

	t.tracks = append(t.tracks, track)

	return len(t.tracks), nil
}

func (t *TrackQueue) pop() (string, error) {
	track := t.tracks[0]
	t.tracks = t.tracks[1:]

	log.Println("Track popped from track queue: ", track)

	return track, nil
}

func (t *TrackQueue) list() []string {
	return t.tracks
}

func (t *TrackQueue) length() int {
	return len(t.tracks)
}
