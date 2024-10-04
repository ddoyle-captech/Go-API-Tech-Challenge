package main

import (
	"Go-API-Tech-Challenge/api/router"
	"Go-API-Tech-Challenge/config"
	"log"
	"net/http"
)

func main() {
	log.Println("... Starting server")

	cfg, err := config.Load(".env.local")
	if err != nil {
		log.Fatalf("Unable to load server config, error: %s", err.Error())
	}

	s := &http.Server{
		Handler: router.New(),
		Addr:    cfg.ServerAddress(),
	}

	log.Printf("Running server at %s ...\n", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("Server crashed and returned error: %s ...\n", err.Error())
	}
}
