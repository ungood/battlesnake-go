package main

import (
	"github.com/ungood/battlesnake-go/server"
)

func main() {
	hostname := "localhost"
	port := 8000

	server.Run(hostname, port)
}
