package battlesnake

import (
	"log"
	"math/rand"

	"github.com/BattlesnakeOfficial/rules/client"
)

// info is called when you create your Battlesnake on play.battlesnake.com
// and controls your Battlesnake's appearance
// TIP: If you open your Battlesnake URL in a browser you should see this data
func Info() client.SnakeMetadataResponse {
	log.Println("INFO")

	return client.SnakeMetadataResponse{
		APIVersion: "1",
		Author:     "ungood",  // TODO: Your Battlesnake username
		Color:      "#ff0000", // TODO: Choose color
		Head:       "default", // TODO: Choose head
		Tail:       "default", // TODO: Choose tail
	}
}

// start is called when your Battlesnake begins a game
func Start(state client.SnakeRequest) {
	log.Println("GAME START")
}

// end is called when your Battlesnake finishes a game
func End(state client.SnakeRequest) {
	log.Printf("GAME OVER\n\n")
}

// move is called on every turn and returns your next move
// Valid moves are "up", "down", "left", or "right"
// See https://docs.battlesnake.com/api/example-move for available data
func Move(state client.SnakeRequest) client.MoveResponse {
	walls := map[client.Coord]bool{}

	for _, snake := range state.Board.Snakes {
		log.Printf("MOVE %d: %s\n", state.Turn, snake.Name)
		walls[snake.Head] = true
		for _, body := range snake.Body {
			walls[body] = true
		}
	}

	isMoveSafe := map[Direction]bool{}

	myHead := state.You.Head

	for direction, coord := range Directions {
		destination := Add(myHead, coord)
		if destination.X < 0 || destination.X >= state.Board.Width {
			isMoveSafe[direction] = false
		} else if destination.Y < 0 || destination.Y >= state.Board.Height {
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
		log.Printf("MOVE %d: No safe moves detected! Moving down\n", state.Turn)
		return client.MoveResponse{Move: "down"}
	}

	// Choose a random move from the safe ones
	nextMove := safeMoves[rand.Intn(len(safeMoves))]

	// TODO: Step 4 - Move towards food instead of random, to regain health and survive longer
	// food := state.Board.Food

	log.Printf("MOVE %d: %s\n", state.Turn, nextMove)
	return client.MoveResponse{
		Move: string(nextMove),
	}
}
