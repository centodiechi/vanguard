package main

import (
	"log"

	"github.com/centodiechi/vanguard/internal"
)

func main() {
	if err := internal.StartServer(); err != nil {
		log.Fatalf("failed to start server %v", err)
	}
}
