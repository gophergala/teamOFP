package main

import (
	"log"
	"time"
)

type Track struct {
	Id       string `json:"id"`
	Time     string `json:"time"`
	Name     string `json:"name"`
	Artist   string `json:"artist"`
	Album    string `json:"album"`
	AlbumArt string `json:"album_art"`
}

type TrackQueue struct {
	tracks    []Track
	updatedAt time.Time
}

func (t *TrackQueue) push(track Track) (int, error) {
	log.Println("Track pushed to track queue: ", track.Name)

	for _, v := range t.tracks {
		if v.Id == track.Id {
			return len(t.tracks), nil
		}
	}
	t.tracks = append(t.tracks, track)

	return len(t.tracks), nil
}

func (t *TrackQueue) pop() (Track, error) {
	track := t.tracks[0]
	t.tracks = t.tracks[1:]

	log.Println("Track popped from track queue: ", track)

	return track, nil
}

func (t *TrackQueue) list() []Track {
	return t.tracks
}

func (t *TrackQueue) length() int {
	return len(t.tracks)
}

func (t *TrackQueue) remove(ID string) {

	for i, v := range t.tracks {
		if v.Id == ID {
			t.tracks = append(t.tracks[:i], t.tracks[i+1:]...)
		}
	}
}
