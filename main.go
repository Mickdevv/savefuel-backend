package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Mickdevv/savefuel-backend/api"
	"github.com/Mickdevv/savefuel-backend/api/documents"
	"github.com/Mickdevv/savefuel-backend/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()

	dbURL := os.Getenv("DB_URL")
	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Error connecting to the database: %s", err)
	}

	serverConfig := api.ServerConfig{
		JWT_SECRET: os.Getenv("JWT_SECRET"),
		DB:         database.New(dbConn),
	}

	mux := http.NewServeMux()

	fsHandler := http.StripPrefix("/static", http.FileServer(http.Dir(os.Getenv("STATIC_FILES_DIR"))))
	mux.Handle("/static/", fsHandler)

	documents.RegisterRoutes(mux, &serverConfig)

	serverPort := os.Getenv("SERVER_PORT")
	server := &http.Server{
		Addr:    ":" + serverPort,
		Handler: mux,
	}
	fmt.Println("Server listening on port", serverPort)
	log.Fatal(server.ListenAndServe())
}
