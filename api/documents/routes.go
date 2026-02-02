package documents

import (
	"net/http"

	"github.com/Mickdevv/savefuel-backend/api"
	"github.com/Mickdevv/savefuel-backend/internal/auth"
)

func RegisterRoutes(mux *http.ServeMux, serverCfg *api.ServerConfig) {
	mux.HandleFunc("GET /documents", api.WithCfg(serverCfg, getDocuments))
	mux.HandleFunc("GET /documents/{id}", api.WithCfg(serverCfg, getDocumentById))
	mux.HandleFunc("PUT /documents/{id}", auth.AuthMiddleware(serverCfg, api.WithCfg(serverCfg, updateDocument)))
	mux.HandleFunc("DELETE /documents/{id}", auth.AuthMiddleware(serverCfg, api.WithCfg(serverCfg, deleteDocument)))
	mux.HandleFunc("POST /documents", auth.AuthMiddleware(serverCfg, api.WithCfg(serverCfg, uploadDocument)))
}
