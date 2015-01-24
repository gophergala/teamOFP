// Spotify Remote API
//
// TeamOFP - GopherGala 2015
//

package main

import (
	"database/sql"
	//"encoding/json"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/crowdmob/goamz/sqs"
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
	sqs *sqs.Queue
}

var context = &Context{}

// GetInfo - Info Endpoint. Returns versioning info.
func GetInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Spotify Remote API v0.1.0")
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

	// Setup App Context
	s, err := sqs.NewFrom(os.Getenv("AWS_ACCESS"), os.Getenv("AWS_SECRET"), "us-east-1")
	if err != nil {
		log.Panic(err)
	}

	q, err := s.GetQueue("spotify-ofp")
	if err != nil {
		log.Panic(err)
	}

	context.sqs = q

	messages := make(chan *sqs.Message)
	go listenOnQueue(context.sqs, messages)
	go processQueue(messages)

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

	//tq := &TrackQueue{}

	// Setup router
	n.UseHandler(r)

	// Start Serve
	if os.Getenv("PORT") != "" {
		n.Run(strings.Join([]string{":", os.Getenv("PORT")}, ""))
	} else {
		n.Run(":8080")
	}

}
