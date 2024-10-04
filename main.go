package main

import (
	"Go-API-Tech-Challenge/api/resources/course"
	"Go-API-Tech-Challenge/api/router"
	"Go-API-Tech-Challenge/config"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("... Starting server")

	cfg, err := config.Load(".env.local")
	if err != nil {
		log.Fatalf("Unable to load server config, error: %s", err.Error())
	}

	db := connectToDB(cfg)

	cr := course.NewRepo(db)

	s := &http.Server{
		Handler: router.New(cr),
		Addr:    cfg.ServerAddress(),
	}

	log.Printf("Running server at %s ...\n", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("Server crashed and returned error: %s ...\n", err.Error())
	}
}

func connectToDB(cfg *config.Config) *sql.DB {
	connection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	// PGX is a DB driver for Postgres.
	db, err := sql.Open("pgx", connection)
	if err != nil {
		log.Fatalf("Unable to open database, error: %s", err.Error())
	}

	// sql.Open() validates our connection info, but doesn't actually connect to the DB. db.Ping()
	// is called here to verify we can talk to the DB
	if err := db.Ping(); err != nil {
		log.Fatalf("Unable to connect to database, error: %s", err.Error())
	}

	return db
}
