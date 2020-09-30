package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	version = "v0.3.5"
	port    = "8080"
)

type Pong struct {
	Respo string
	Stat  string
}

//GitCommit get commitid from git
var GitCommit string

// AppVersion - Application verison constant
var fullVersion = version + "-" + GitCommit

func main() {

	fmt.Println("Starting RuleBase Engine Version:" + fullVersion + " on port " + port)

	http.HandleFunc("/ping", ping())
	fmt.Println("starting to listen on port " + port)
	http.ListenAndServe(":"+port, nil)
}

func ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("Activated Ping test")
		responseforping := Pong{"pong", "ok"}
		js, err := json.Marshal(responseforping)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}
