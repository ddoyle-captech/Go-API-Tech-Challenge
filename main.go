package main

import (
	"Go-API-Tech-Challenge/api/resources/course"
	"Go-API-Tech-Challenge/api/router"
	"Go-API-Tech-Challenge/config"
	"database/sql"
	"log"
	"net/http"
)

func main() {
	log.Println("... Starting server")

	// TODO: add proper DB connection w/ driver + connection info
	cr := course.NewRepo(&sql.DB{})

	cfg, err := config.Load(".env.local")
	if err != nil {
		log.Fatalf("Unable to load server config, error: %s", err.Error())
	}

	s := &http.Server{
		Handler: router.New(cr),
		Addr:    cfg.ServerAddress(),
	}

	log.Printf("Running server at %s ...\n", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("Server crashed and returned error: %s ...\n", err.Error())
	}
}
