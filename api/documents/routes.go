package documents

import (
	"github.com/Mickdevv/savefuel-backend/api"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, serverCfg *api.ServerConfig) {
	mux.HandleFunc("GET /documents", api.WithCfg(serverCfg, getDocuments))
	mux.HandleFunc("GET /documents/{id}", api.WithCfg(serverCfg, getDocumentById))
	mux.HandleFunc("PUT /documents/{id}", api.WithCfg(serverCfg, updateDocument))
	mux.HandleFunc("DELETE /documents/{id}", api.WithCfg(serverCfg, deleteDocument))
	mux.HandleFunc("POST /documents", api.WithCfg(serverCfg, uploadDocument))
}
