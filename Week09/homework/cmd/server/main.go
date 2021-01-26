package main

import (
	"homework/internal/server"
	"log"
)

func main() {
	if err := server.StartServer(); err != nil {
		log.Fatalf("err: %v", err)
	}
}
