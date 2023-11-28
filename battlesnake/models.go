package battlesnake

import (
	. "github.com/BattlesnakeOfficial/rules/client"
)

func Add(a, b Coord) Coord {
	return Coord{
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

var Directions = map[Direction]Coord{
	Up:    {X: 0, Y: 1},
	Down:  {X: 0, Y: -1},
	Left:  {X: -1, Y: 0},
	Right: {X: 1, Y: 0},
}
