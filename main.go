package main

import (
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/twilio/twilio-go"
)

type SessionTracker struct {
	ClientMSGS   []string
	ClientID     string
	ClientNumber string
}

type Server struct {
	sessions map[string]*SessionTracker
	mut      sync.Mutex
}

func main() {
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("port env not provided :c")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	server := &Server{
		sessions: make(map[string]*SessionTracker),
	}

	router.Get("/ready", handlerReady)
	router.Get("/err", handlerError)

	router.Post("/bot", server.handlerConvo)

	serve := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port http://localhost:%v", portString)

	err := serve.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
