package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/riju-stone/sevin/api/models"
	"github.com/riju-stone/sevin/api/services"
	"github.com/riju-stone/sevin/api/utils"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	utils.InitCustomLogger()
	l := utils.CustomLogger

	// Connect to rabbitmq
	rabbitmqClient, err := services.ConnectToRabbitMQ()
	if err != nil {
		l.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	rabbitmqClient.InitTaskQueue()

	// Connect to database
	db, err := services.ConnectToDB()
	if err != nil {
		l.Fatalf("Failed to connect to the database: %v", err)
	}

	// Initialize the database tables
	db.AutoMigrate(&models.Task{})

	// Initialize the router
	r := mux.NewRouter()
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
		l.Info("Health check endpoint hit")
	})

	http.Handle("/", r)

	// Start the server
	l.Info("Starting server on port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		l.Panicf("Error starting server: %v", err)
	}
}
