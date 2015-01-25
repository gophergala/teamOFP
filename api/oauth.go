package main

import (
	"encoding/json"
	"golang.org/x/oauth2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type User struct {
	UserID      string    `db:"user_id"`
	AccessToken string    `db:"access_token"`
	AvatarURL   string    `db:"avatar_url"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func initOauth2() error {

	conf := &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Scopes:       []string{""},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
		},
	}

	context.oauth = conf

	return nil
}

func getRedirectURL() string {
	url := context.oauth.AuthCodeURL("state", oauth2.AccessTypeOffline)

	log.Println("Redirect URL: ", url)

	return url
}

func createUser(token string) (string, error) {
	// TODO:
	// - Get github user ID
	// - Store in DB

	url := strings.Join([]string{"https://api.github.com/user?", "access_token=", token}, "")
	resp, err := http.Get(url)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var ghUser map[string]interface{}
	err = json.Unmarshal(body, &ghUser)
	if err != nil {
		log.Panic(err)
	}

	// Insert into DB

	log.Println("User: ", ghUser)

	userID := int(ghUser["id"].(float64))

	id := strconv.Itoa(userID)

	return id, nil
}

// Auth - Endpoint
func Auth(w http.ResponseWriter, r *http.Request) {
	initOauth2()
	url := getRedirectURL()

	http.Redirect(w, r, url, http.StatusUnauthorized)
}

// Callback - Endpoint
// Auth
func Callback(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")

	log.Println("Code: ", code)

	// Swap temp token
	tok, err := context.oauth.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Panic(err)
	}
	log.Println("Token: ", tok)

	id, _ := createUser(tok.AccessToken)

	// Create session
	session, _ := store.Get(r, "groupify")

	log.Println("Saving into session: ", id)

	// Set some session values.
	session.Values["userID"] = id

	// Save it.
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusAccepted)
}
