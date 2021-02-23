package main

import (
	"log"

	"github.com/jorgeAM/goHexagonal/internal/platform/server"
)

const (
	host = "localhost"
	port = 3000
)

func main() {
	srv := server.NewServer(host, port)

	if err := srv.Run(); err != nil {
		log.Fatalf("something got wrong %v", err)
	}
}
