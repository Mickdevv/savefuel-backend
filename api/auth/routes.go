package auth

import (
	"net/http"

	"github.com/Mickdevv/savefuel-backend/api"
	"github.com/Mickdevv/savefuel-backend/internal/auth"
)

func RegisterRoutes(mux *http.ServeMux, serverCfg *api.ServerConfig) {
	mux.HandleFunc("POST /register", api.WithCfg(serverCfg, registerHandler))
	mux.HandleFunc("POST /login", api.WithCfg(serverCfg, loginHandler))
	mux.HandleFunc("POST /refresh", auth.AuthMiddleware(serverCfg, api.WithCfg(serverCfg, refreshTokenHandler)))
}
