package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/BattlesnakeOfficial/rules/client"
	"github.com/ungood/battlesnake-go/actor"
)

var snake = &actor.SimpleSnake{}

func decode(r *http.Request, event *zerolog.Event, request interface{}) {
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		log.Err(err).Msg("Failed to decode request")
	}

	event.Any("request", request)
}

func encode(w http.ResponseWriter, event *zerolog.Event, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Err(err).Msg("Failed to encode response")
	}

	event.Any("response", response)
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	event := log.Info()

	response := snake.Info()
	encode(w, event, response)

	event.Msg("Handled Info request")
}

func HandleStart(w http.ResponseWriter, r *http.Request) {
	event := log.Info()

	request := client.SnakeRequest{}
	decode(r, event, &request)

	snake.Start(request)

	event.Msg("Handled Start request")
}

func HandleMove(w http.ResponseWriter, r *http.Request) {
	event := log.Info()

	request := client.SnakeRequest{}
	decode(r, event, &request)

	response := snake.Move(request)
	encode(w, event, response)

	event.Msg("Handled Move request")
}

func HandleEnd(w http.ResponseWriter, r *http.Request) {
	event := log.Info()

	request := client.SnakeRequest{}
	decode(r, event, &request)

	snake.End(request)

	event.Msg("Handled End request")
}

const ServerID = "ungood/battlesnake-go"

func withServerID(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Server", ServerID)
		next(w, r)
	}
}

// Start Battlesnake Server
// TODO: Figure out hosting
// TODO: Add paths that initialize snakes with different strategies
func Run(hostname string, port int) {
	http.HandleFunc("/", withServerID(HandleIndex))
	http.HandleFunc("/start", withServerID(HandleStart))
	http.HandleFunc("/move", withServerID(HandleMove))
	http.HandleFunc("/end", withServerID(HandleEnd))

	addr := fmt.Sprintf("%s:%d", hostname, port)
	url := fmt.Sprintf("http://%s", addr)

	log.Info().
		Str("addr", addr).
		Str("url", url).
		Msg("Starting server")

	err := http.ListenAndServe(addr, nil)
	log.Fatal().Err(err).Msg("Failed to start server")
}
