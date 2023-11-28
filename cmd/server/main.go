package main

import (
	"github.com/ungood/battlesnake-go/battlesnake"
)

func main() {
	hostname := "localhost"
	port := 8000

	battlesnake.Run(hostname, port)
}
