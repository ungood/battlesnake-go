package actor

import (
	"github.com/BattlesnakeOfficial/rules/client"
)

type SnakeActor interface {
	// Info is called when you create your Battlesnake on play.battlesnake.com
	Info() client.SnakeMetadataResponse

	// Start is called when your Battlesnake begins a game
	Start(request client.SnakeRequest)

	// End is called when your Battlesnake finishes a game
	End(request client.SnakeRequest)

	// Move is called on every turn and returns your next move
	// See https://docs.battlesnake.com/api/example-move for available data
	Move(request client.SnakeRequest) client.MoveResponse
}

var Actors = map[string]SnakeActor{
	"simple": &SimpleSnake{},
}
