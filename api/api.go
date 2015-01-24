// Spotify Remote API
//
// TeamOFP - GopherGala 2015
//

package main

import (
	"database/sql"
	//"encoding/json"
	"encoding/json"
	"fmt"
	"github.com/codegangsta/negroni"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strings"
)

// App context
type Context struct {
	db *sql.DB
	//airbrake *gobrake.Notifier
}

type Track struct {
	Id     string `json:"id"`
	Time   string `json:"time"`
	Name   string `json:"name"`
	Artist string `json:"artist"`
	Album  string `json:"album"`
}

var context = &Context{}

// GetInfo - Info Endpoint. Returns versioning info.
func GetInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Spotify Remote API v0.1.0")
}

func SearchSpotify(w http.ResponseWriter, r *http.Request) {
	response, _ := json.Marshal(tracks)
	w.Write(response)
}

var tracks []Track

func init() {
	s1 := Track{Time: "3:07", Name: "The Ride", Artist: "David Allan Cole", Album: "16 Biggest Hits"}
	s2 := Track{Time: "1:20", Name: "Bookends", Artist: "Simon and Garfunkel", Album: "Greatest Hits"}
	s3 := Track{Time: "3:28", Name: "A Woman Left Lonely", Artist: "Janis Joplin", Album: "The Pearl Sessions"}
	tracks = []Track{s1, s2, s3}
}

func main() {
	log.Println("Starting Spotify Remote API...")

	// Load .env
	err := godotenv.Load()
	if err != nil {
		// Can't load .env, so setenv defaults
		os.Setenv("SQL_HOST", "localhost:8091")
		os.Setenv("SQL_USER", "root")
		os.Setenv("SQL_PASSWORD", "")
		os.Setenv("SQL_DB", "spotify_remote")
	}

	router := mux.NewRouter()
	r := router.PathPrefix("/api/v1").Subrouter() // Prepend API Version

	// Setup Negroni
	n := negroni.Classic()

	// Info
	r.HandleFunc("/", GetInfo).Methods("GET")

	// TrackQueue
	r.HandleFunc("/queue/add", PostAddTrack).Methods("POST")
	r.HandleFunc("/queue/list", GetListTracks).Methods("GET")
	//r.HandleFunc("/queue/upvote", AddTrack).Methods("POST")
	//r.HandleFunc("/queue/downvote", AddTrack).Methods("POST")

	r.HandleFunc("/search", SearchSpotify).Methods("GET")

	tq := &TrackQueue{}

	trackList := tq.list()
	log.Println("Track Queue: ", trackList)

	tq.push("song1")
	tq.push("song2")

	trackList = tq.list()
	log.Println("Track Queue: ", trackList)

	track, _ := tq.pop()
	log.Println("Track: ", track)

	trackList = tq.list()
	log.Println("Track Queue: ", trackList)

	log.Println("Track Queue Length: ", tq.length())

	// Setup router
	n.UseHandler(r)

	// Start Serve
	if os.Getenv("PORT") != "" {
		n.Run(strings.Join([]string{":", os.Getenv("PORT")}, ""))
	} else {
		n.Run(":8080")
	}

}
