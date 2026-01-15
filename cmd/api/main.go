package main

import (
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"
	"github.com/vfaust1/movie-api/internal/store"
)

func main() {
	
	if err := godotenv.Load(); err != nil {
		log.Println("Info: No .env file found")
	}

	if err := store.InitDB(); err != nil {
		log.Fatal("Failed to initialize DB: ", err)
	}

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      routes(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  1 * time.Minute,
	}

	log.Println("🎬 Server started on http://localhost:8080")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}