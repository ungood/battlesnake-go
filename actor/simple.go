package actor

import (
	"math/rand"

	"github.com/rs/zerolog/log"

	"github.com/BattlesnakeOfficial/rules/client"
)

type SimpleSnake struct{}

func (s *SimpleSnake) Info() client.SnakeMetadataResponse {
	return client.SnakeMetadataResponse{
		APIVersion: "1",
		Author:     "ungood",     // TODO: Your Battlesnake username
		Color:      "#ff00ff",    // TODO: Choose color
		Head:       "tiger-king", // TODO: Choose head
		Tail:       "tiger-tail", // TODO: Choose tail
	}
}

func (s *SimpleSnake) Start(request client.SnakeRequest) {}

func (s *SimpleSnake) End(request client.SnakeRequest) {}

func (s *SimpleSnake) Move(request client.SnakeRequest) client.MoveResponse {
	walls := map[client.Coord]bool{}

	for _, snake := range request.Board.Snakes {
		walls[snake.Head] = true
		for _, body := range snake.Body {
			walls[body] = true
		}
	}

	isMoveSafe := map[Direction]bool{}

	myHead := request.You.Head

	for direction, coord := range Directions {
		destination := Add(myHead, coord)
		if destination.X < 0 || destination.X >= request.Board.Width {
			isMoveSafe[direction] = false
		} else if destination.Y < 0 || destination.Y >= request.Board.Height {
			isMoveSafe[direction] = false
		} else if walls[destination] {
			isMoveSafe[direction] = false
		} else {
			isMoveSafe[direction] = true
		}
	}

	// TODO: Step 1 - Prevent your Battlesnake from moving out of bounds
	// DONE

	// TODO: Step 2 - Prevent your Battlesnake from colliding with itself
	// DONE

	// TODO: Step 3 - Prevent your Battlesnake from colliding with other Battlesnakes
	// DONE

	// TODO: Avoid going to a spot that another snake may move to
	// TODO: The tail of a snake is a safe move!

	// Are there any safe moves left?
	safeMoves := []Direction{}
	for move, isSafe := range isMoveSafe {
		if isSafe {
			safeMoves = append(safeMoves, move)
		}
	}

	if len(safeMoves) == 0 {
		log.Warn().Msg("No safe moves detected!")
		return client.MoveResponse{Move: "down"}
	}

	// Choose a random move from the safe ones
	nextMove := safeMoves[rand.Intn(len(safeMoves))]

	// TODO: Step 4 - Move towards food instead of random, to regain health and survive longer
	// food := state.Board.Food

	return client.MoveResponse{
		Move: string(nextMove),
	}
}
