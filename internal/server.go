package internal

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/centodiechi/vanguard/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func StartServer() error {

	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Use(cors.Default())

	routes.RegisterAuthRoutes(router)

	if err := godotenv.Load(".env"); err != nil {
		log.Println("Error in loading env file %w", err)
	}

	port := os.Getenv("PORT")

	vgServer := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	go func() {
		log.Printf("Starting Vanguard Gateway on port %s\n", port)
		if err := vgServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Panic shut down")
	return nil
}
