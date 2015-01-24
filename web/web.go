package main

import (
	"github.com/codegangsta/negroni"
	"net/http"
	"os"
)

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
}

func queueAction(w http.ResponseWriter, req *http.Request) {
}

func enqueueAction(w http.ResponseWriter, req *http.Request) {
}

func playPauseAction(w http.ResponseWriter, req *http.Request) {
}

func setVolumeAction(w http.ResponseWriter, req *http.Request) {
}
