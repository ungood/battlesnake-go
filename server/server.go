package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/justinas/alice"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"

	"github.com/BattlesnakeOfficial/rules/client"
	"github.com/ungood/battlesnake-go/actor"
)

var snake = &actor.SimpleSnake{}

var arrows = map[string]string{
	"up":    "↑",
	"down":  "↓",
	"left":  "←",
	"right": "→",
}

func decode[T any](r *http.Request, event *zerolog.Event, request *T) {
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		log.Err(err).Msg("Failed to decode request")
	}

	event.Any("request", request)
}

func encode[T any](w http.ResponseWriter, event *zerolog.Event, response *T) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Err(err).Msg("Failed to encode response")
	}

	event.Any("response", response)
}

// Add game info to the logger context
func withGameInfo(r *http.Request, request *client.SnakeRequest) context.Context {
	ctx := r.Context()

	return hlog.FromRequest(r).With().
		Str("game", request.Game.ID).
		Int("turn", request.Turn).
		Logger().WithContext(ctx)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	event := log.Info()

	response := snake.Info(r.Context())
	encode(w, event, &response)

	event.Msg("Handled Info request")
}

func handleStart(w http.ResponseWriter, r *http.Request) {
	event := log.Info()

	request := client.SnakeRequest{}
	decode(r, event, &request)

	ctx := withGameInfo(r, &request)
	snake.Start(ctx, request)

	event.Msgf("Game start")
}

func handleMove(w http.ResponseWriter, r *http.Request) {
	event := log.Info()

	request := client.SnakeRequest{}
	decode(r, event, &request)

	ctx := withGameInfo(r, &request)
	response := snake.Move(ctx, request)
	encode(w, event, &response)

	arrow := arrows[response.Move]
	event.Msgf("Turn %3d: %s", request.Turn, arrow)
}

func handleEnd(w http.ResponseWriter, r *http.Request) {
	event := log.Info()

	request := client.SnakeRequest{}
	decode(r, event, &request)

	ctx := withGameInfo(r, &request)
	snake.End(ctx, request)

	event.Msgf("Game end")
}

// Start Battlesnake Server
// TODO: Add paths that initialize snakes with different strategies
func Run(hostname string, port int) {
	logHandler := hlog.NewHandler(log.Logger)
	var chain = alice.New(logHandler)

	handler := http.NewServeMux()
	handler.Handle("/", chain.ThenFunc(handleIndex))
	handler.Handle("/start", chain.ThenFunc(handleStart))
	handler.Handle("/move", chain.ThenFunc(handleMove))
	handler.Handle("/end", chain.ThenFunc(handleEnd))

	addr := fmt.Sprintf("%s:%d", hostname, port)

	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	log.Info().Msgf("Starting server: http://%s", addr)
	err := server.ListenAndServe()
	log.Fatal().Err(err).Msg("Failed to start server")
}
