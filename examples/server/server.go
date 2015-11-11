package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Path[1:]
	log.Printf("params: %s", params)
	fmt.Fprintf(w, "Welcome, %s!", params)
}

type About struct {
	Timestamp string `json:"timestamp"`
	Message   string `json:"message"`
}

func jsonHandle(w http.ResponseWriter, r *http.Request) {
	about := About{
		Timestamp: time.Now().Format(time.RFC3339),
		Message:   "Golang Brasil",
	}

	log.Printf("about : %v", about)
	if err := json.NewEncoder(w).Encode(about); err != nil {
		panic(err)
	}
}

func main() {
	log.Print("running...")
	http.HandleFunc("/", handler)
	http.HandleFunc("/meetup/about", jsonHandle)
	http.ListenAndServe(":8080", nil)
}
