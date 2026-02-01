package test

import (
	"database/sql"
	"log"
	"os"

	"github.com/Mickdevv/savefuel-backend/api"
	"github.com/Mickdevv/savefuel-backend/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func TestServerCFG() api.ServerConfig {

	godotenv.Load()

	dbURL := os.Getenv("DB_URL")
	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err)
	}

	serverConfig := api.ServerConfig{
		JWT_SECRET:       os.Getenv("JWT_SECRET"),
		DB:               database.New(dbConn),
		STATIC_FILES_DIR: os.Getenv("STATIC_FILES_DIR"),
	}
	return serverConfig
}
