package main

import (
	"log"
	"time"
)

type Track struct {
	Id       string `json:"id" db:"track_id"`
	Time     string `json:"time" db:"time"`
	Name     string `json:"name" db:"name"`
	Artist   string `json:"artist" db:"artist"`
	Album    string `json:"album" db:"album"`
	AlbumArt string `json:"album_art" db:"album_art"`
}

type TrackQueue struct {
	tracks    []Track
	updatedAt time.Time
}

func (t *TrackQueue) push(track Track) (int, error) {
	log.Println("Track pushed to track queue: ", track.Name)

	_, err := context.db.NamedExec("INSERT INTO track_queue (track_id, time, name, artist, album, album_art) VALUES (:track_id, :time, :name, :artist, :album, :album_art)", &track)

	if err != nil {
		return 0, err
	}
	return t.length(), nil
}

func (t *TrackQueue) pop() (Track, error) {

	track := Track{}

	err := context.db.Get(&track, "SELECT track_id, name, artist, album, album_art, time FROM track_queue ORDER BY id ASC LIMIT 1;")
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			// No tracks left in the queue
		} else {
			log.Panic(err)
		}
	}

	if track.Id != "" {
		context.db.Exec("DELETE FROM track_queue WHERE track_id = ?", track.Id)
		log.Println("Track popped from track queue: ", track)
	}
	return track, nil

}

func (t *TrackQueue) list() []Track {
	tq := []Track{}

	err := context.db.Select(&tq, "SELECT track_id, name, artist, album, album_art, time FROM track_queue")
	if err != nil {
		log.Panic(err)
	}

	return tq
}

func (t *TrackQueue) length() int {
	var count int
	err := context.db.Get(&count, "SELECT count(*) FROM track_queue")
	if err != nil {
		log.Panic(err)
	}

	return count
}

func (t *TrackQueue) remove(ID string) {

	context.db.Exec("DELETE FROM track_queue WHERE track_id = ?", ID)
}
