package main

import "github.com/free-go-foundation/go-engine/server"

func main() {
	s := server.NewServer(":3001")
	s.Run()
}
