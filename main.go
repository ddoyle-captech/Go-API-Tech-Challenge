package main

import (
	"Go-API-Tech-Challenge/api/router"
	"log"
	"net/http"
)

func main() {
	log.Println("... Starting server")

	s := &http.Server{
		Handler: router.New(),
		Addr:    ":8000",
	}

	log.Printf("Running server at %s ...\n", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("Server crashed and returned error: %s ...\n", err.Error())
	}
}
