package main

import (
	"fmt"
	//"log"
	"encoding/json"
	"net/http"
)

// PostAddTrack - Add track to Track Queue
// Format (JSON): {"<track_id>"}
func PostAddTrack(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "140 Proof Persona API v0.1.0")
}

// GetListTracks - Retrieve list of tracks in Track Queue
func GetListTracks(w http.ResponseWriter, r *http.Request) {
	response, _ := json.Marshal(tracks)
	w.Write(response)
}
