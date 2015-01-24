package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
)

// start of script tag
const ScriptStart = "tell application \"Spotify\" to "

var commands = map[string]string{
	"state":      "state",
	"play":       "play",
	"pause":      "pause",
	"stop":       "stop",
	"duration":   "duration of current track",
	"name":       "name of current track",
	"album":      "album of current track",
	"id":         "id of current track",
	"artwork":    "artwork of current track",
	"vol_loud":   "set sound volume to 100",
	"vol_soft":   "set sound volume to 20",
	"vol_norm":   "set sound volume to 50",
	"set_volume": "set sound volume to ", //requires parameter
	"play_track": "play track ",          //requires parameter
}

func createSystemCall(command string, param string) string {
	return ScriptStart + commands[command] + param
}

func main() {
	var cmd = flag.String("o", "pause", "Enter the command for spotify")
	flag.Parse()

	command := createSystemCall(*cmd, "")
	if command == "" {
		fmt.Println("exiting...")
	}
	fmt.Println(command)
	out, err := exec.Command("/usr/bin/osascript", "-e", command).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The output is %s\n", out)
}
