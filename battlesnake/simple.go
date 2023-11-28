package battlesnake

import (
	"log"
	"math/rand"

	"github.com/BattlesnakeOfficial/rules/client"
)

type SimpleSnake struct{}

func (s *SimpleSnake) Info() client.SnakeMetadataResponse {
	log.Println("INFO")

	return client.SnakeMetadataResponse{
		APIVersion: "1",
		Author:     "ungood",  // TODO: Your Battlesnake username
		Color:      "#ff0000", // TODO: Choose color
		Head:       "default", // TODO: Choose head
		Tail:       "default", // TODO: Choose tail
	}
}

func (s *SimpleSnake) Start(request client.SnakeRequest) {
	log.Println("GAME START")
}

func (s *SimpleSnake) End(request client.SnakeRequest) {
	log.Printf("GAME OVER\n\n")
}

func (s *SimpleSnake) Move(request client.SnakeRequest) client.MoveResponse {
	walls := map[client.Coord]bool{}

	for _, snake := range request.Board.Snakes {
		log.Printf("MOVE %d: %s\n", request.Turn, snake.Name)
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

	// Are there any safe moves left?
	safeMoves := []Direction{}
	for move, isSafe := range isMoveSafe {
		if isSafe {
			safeMoves = append(safeMoves, move)
		}
	}

	if len(safeMoves) == 0 {
		log.Printf("MOVE %d: No safe moves detected! Moving down\n", request.Turn)
		return client.MoveResponse{Move: "down"}
	}

	// Choose a random move from the safe ones
	nextMove := safeMoves[rand.Intn(len(safeMoves))]

	// TODO: Step 4 - Move towards food instead of random, to regain health and survive longer
	// food := state.Board.Food

	log.Printf("MOVE %d: %s\n", request.Turn, nextMove)
	return client.MoveResponse{
		Move: string(nextMove),
	}
}
