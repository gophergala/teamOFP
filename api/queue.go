package main

import (
	"encoding/json"
	//"fmt"
	"log"
	"net/http"
)

type remoteCommand struct {
	Command string `json:"command"`
	Param   string `json:"param"`
}

type response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type nowPlaying struct {
	current       Track `json:"current"`
	timeRemaining int   `json:"time_running"`
}

func queueNextTrack() error {
	// Received track ended, so send next track from track queue

	nextTrack, _ := context.tq.pop()
	queueTrackRemote(nextTrack.Id)

	return nil
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

func updateNowPlayingTrack(ID string) {

	t := getTrackDetails(ID)

	context.np.current = *t
}

func updateNowPlayingTime(time int) {
	context.np.timeRemaining = time
}

// PostAddTrack - Add track to Track Queue
// Format (JSON): {"<track_id>"}
func PostAddTrack(w http.ResponseWriter, r *http.Request) {

	reqData := map[string]string{}

	// Parse JSON Data
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&reqData)
	if err != nil {
		log.Fatal(err)
	}

	// Get track details
	t := getTrackDetails(reqData["track_id"])

	context.tq.push(*t)

	resp := response{
		Status:  "success",
		Message: "Track Added to Queue: " + t.Id,
	}

	jresp, _ := json.Marshal(resp)

	w.Write(jresp)
}

// GetListTracks - Retrieve list of tracks in Track Queue
func GetListTracks(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
	queuedTracks := context.tq.list()

	response["queue"] = queuedTracks
	response["now_playing"] = map[string]interface{}{
		"track":          context.np.current,
		"time_remaining": context.np.timeRemaining,
	}

	jresponse, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jresponse)
}

// DeleteTrack - Delete a track from the Track Queue
func PostDeleteTrack(w http.ResponseWriter, r *http.Request) {

	reqData := map[string]string{}

	// Parse JSON Data
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&reqData)
	if err != nil {
		log.Fatal(err)
	}

	t := reqData["track_id"]

	context.tq.remove(t)

	w.WriteHeader(204)
	w.Write([]byte(`{"status":"deleted", "track":"` + t + `"}`))
}
