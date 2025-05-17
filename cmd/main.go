package main

import (
	"log"

	"github.com/centodiechi/vanguard/internal"
	syncprovider "github.com/centodiechi/vanguard/sync-provider"
)

func init() {
	if err := syncprovider.Initialize(); err != nil {
		log.Fatalf("failed to initialize sync provider %v", err)
	}
}

func main() {
	if err := internal.StartServer(); err != nil {
		log.Fatalf("failed to start server %v", err)
	}
}
