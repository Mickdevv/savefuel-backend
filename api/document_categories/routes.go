package document_categories

import (
	"net/http"

	"github.com/Mickdevv/savefuel-backend/api"
	"github.com/Mickdevv/savefuel-backend/internal/auth"
)

func RegisterRoutes(mux *http.ServeMux, serverCfg *api.ServerConfig) {
	mux.HandleFunc("GET /document_categories", getDocuments(serverCfg))
	mux.HandleFunc("GET /document_categories/{id}", getDocumentById(serverCfg))
	mux.HandleFunc("POST /document_categories", auth.AuthMiddleware(serverCfg, createDocumentCategory(serverCfg)))
	mux.HandleFunc("PUT /document_categories/{id}", auth.AuthMiddleware(serverCfg, updateDocumentCategory(serverCfg)))
	mux.HandleFunc("DELETE /document_categories/{id}", auth.AuthMiddleware(serverCfg, deleteDocumentCategory(serverCfg)))
}
