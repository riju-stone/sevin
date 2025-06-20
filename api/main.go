package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
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
	services.InitTaskQueue()

	// Connect to database
	services.InitDB()

	r := mux.NewRouter()

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
		l.Info("Health check endpoint hit")
	})

	http.Handle("/", r)
	l.Info("Starting server on port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		l.Panicf("Error starting server: %v", err)
	}
}
