package testUtils

import (
	"database/sql"
	"log"
	"net/url"
	"os"

	"github.com/Mickdevv/savefuel-backend/api"
	"github.com/Mickdevv/savefuel-backend/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func TestServerCFG() api.ServerConfig {

	godotenv.Load("../../.env")

	dbURL := os.Getenv("DB_URL")

	if dbURL == "" {
		log.Fatal("Test DB url is empty")
	}

	u, err := url.Parse(dbURL)
	if err != nil {
		log.Fatal(err)
	}

	q := u.Query()
	q.Set("sslmode", "disable")
	u.RawQuery = q.Encode()

	dbConn, err := sql.Open("postgres", u.String())
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
