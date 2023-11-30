package actor

import (
	client "github.com/BattlesnakeOfficial/rules/client"
)

func Add(a, b client.Coord) client.Coord {
	return client.Coord{
		X: a.X + b.X,
		Y: a.Y + b.Y,
	}
}

type Direction string

const (
	Up    Direction = "up"
	Down  Direction = "down"
	Left  Direction = "left"
	Right Direction = "right"
)

var Directions = map[Direction]client.Coord{
	Up:    {X: 0, Y: 1},
	Down:  {X: 0, Y: -1},
	Left:  {X: -1, Y: 0},
	Right: {X: 1, Y: 0},
}
