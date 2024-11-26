package main

import (
	"context"
	"fmt"
	"log"
	"marrywith-gin-api/config"
	"marrywith-gin-api/controllers"
	"marrywith-gin-api/utils"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize Gin router
	router := gin.Default()

	// Initialize MongoDB client
	client := utils.ConnectMongoDB(cfg.MongoURI)
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatalf("Error disconnecting MongoDB: %v", err)
		}
	}()

	// Check the connection
	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatalf("Could not connect to MongoDB: %v", err)
	}
	fmt.Println("Connected to MongoDB!")

	// Initialize controllers
	personController := controllers.NewPersonController(client)

	// Define routes
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	router.POST("/person", personController.CreatePerson)
	router.GET("/persons", personController.GetPersons)

	// Set up graceful shutdown
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Gracefully close the server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
