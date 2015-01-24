package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type remoteCommand struct {
	Command string `json:"command"`
	Param   string `json:"param"`
}

// PostAddTrack - Add track to Track Queue
// Format (JSON): {"<track_id>"}
func PostAddTrack(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "140 Proof Persona API v0.1.0")

	context.tq.push("track123")
}

// GetListTracks - Retrieve list of tracks in Track Queue
func GetListTracks(w http.ResponseWriter, r *http.Request) {
	response, _ := json.Marshal(tracks)
	w.Write(response)
}

// queueTrackRemote - Queues a track remotely
func queueTrackRemote(track string) {

	m := remoteCommand{
		Command: "play_track",
		Param:   track,
	}

	err := pushMessage(context.sqs, m)
	if err != nil {
		log.Panic(err)
	}

	log.Println("Track Queued: ", track)
}
