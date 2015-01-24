package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
)

type Track struct {
	Id     string `json:"id"`
	Time   string `json:"time"`
	Name   string `json:"name"`
	Artist string `json:"artist"`
	Album  string `json:"album"`
}

var tracks []Track

func init() {
	s1 := Track{Time: "3:07", Name: "The Ride", Artist: "David Allan Cole", Album: "16 Biggest Hits"}
	s2 := Track{Time: "1:20", Name: "Bookends", Artist: "Simon and Garfunkel", Album: "Greatest Hits"}
	s3 := Track{Time: "3:28", Name: "A Woman Left Lonely", Artist: "Janis Joplin", Album: "The Pearl Sessions"}
	tracks = []Track{s1, s2, s3}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/search.json", searchAction)
	mux.HandleFunc("/enqueue.json", enqueueAction)
	mux.HandleFunc("/queue.json", queueAction)
	mux.HandleFunc("/play_pause.json", playPauseAction)
	mux.HandleFunc("/volume.json", setVolumeAction)

	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run("localhost:" + os.Getenv("PORT"))
}

func searchAction(w http.ResponseWriter, req *http.Request) {
	response, _ := json.Marshal(tracks)
	w.Write(response)
}

func queueAction(w http.ResponseWriter, req *http.Request) {
	response, _ := json.Marshal(tracks)
	w.Write(response)
}

func enqueueAction(w http.ResponseWriter, req *http.Request) {
}

func playPauseAction(w http.ResponseWriter, req *http.Request) {
}

func setVolumeAction(w http.ResponseWriter, req *http.Request) {
}
