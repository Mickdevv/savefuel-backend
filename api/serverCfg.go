package api

import (
	"net/http"

	"github.com/Mickdevv/savefuel-backend/internal/database"
)

type ServerConfig struct {
	JWT_SECRET       string
	DB               *database.Queries
	STATIC_FILES_DIR string
}

type AppHandler func(cfg *ServerConfig, w http.ResponseWriter, r *http.Request)

func WithCfg(cfg *ServerConfig, handler AppHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(cfg, w, r)
	}
}
