package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/BattlesnakeOfficial/rules/client"
	"github.com/ungood/battlesnake-go/pkg/battlesnake"
)

// HTTP Handlers

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	response := battlesnake.Info()

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("ERROR: Failed to encode info response, %s", err)
	}
}

func HandleStart(w http.ResponseWriter, r *http.Request) {
	state := client.SnakeRequest{}
	err := json.NewDecoder(r.Body).Decode(&state)
	if err != nil {
		log.Printf("ERROR: Failed to decode start json, %s", err)
		return
	}

	battlesnake.Start(state)

	// Nothing to respond with here
}

func HandleMove(w http.ResponseWriter, r *http.Request) {
	state := client.SnakeRequest{}
	err := json.NewDecoder(r.Body).Decode(&state)
	if err != nil {
		log.Printf("ERROR: Failed to decode move json, %s", err)
		return
	}

	response := battlesnake.Move(state)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("ERROR: Failed to encode move response, %s", err)
		return
	}
}

func HandleEnd(w http.ResponseWriter, r *http.Request) {
	state := client.SnakeRequest{}
	err := json.NewDecoder(r.Body).Decode(&state)
	if err != nil {
		log.Printf("ERROR: Failed to decode end json, %s", err)
		return
	}

	battlesnake.End(state)

	// Nothing to respond with here
}

// Middleware

const ServerID = "ungood/battlesnake-go"

func withServerID(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Server", ServerID)
		next(w, r)
	}
}

// Start Battlesnake Server
func Run(hostname string, port int) {
	http.HandleFunc("/", withServerID(HandleIndex))
	http.HandleFunc("/start", withServerID(HandleStart))
	http.HandleFunc("/move", withServerID(HandleMove))
	http.HandleFunc("/end", withServerID(HandleEnd))

	addr := fmt.Sprintf("%s:%d", hostname, port)

	log.Printf("Running Battlesnake at %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
