package auth

import (
	"net/http"

	"github.com/Mickdevv/savefuel-backend/api"
	"github.com/Mickdevv/savefuel-backend/internal/auth_helpers"
)

func RegisterRoutes(mux *http.ServeMux, serverCfg *api.ServerConfig) {
	mux.HandleFunc("POST /register", api.WithCfg(serverCfg, RegisterHandler))
	mux.HandleFunc("POST /login", api.WithCfg(serverCfg, LoginHandler))
	mux.HandleFunc("POST /refresh", auth_helpers.AuthMiddleware(serverCfg, api.WithCfg(serverCfg, RefreshTokenHandler)))
}
