package document_categories

import (
	"net/http"

	"github.com/Mickdevv/savefuel-backend/api"
	"github.com/Mickdevv/savefuel-backend/internal/auth_helpers"
)

func RegisterRoutes(mux *http.ServeMux, serverCfg *api.ServerConfig) {
	mux.HandleFunc("GET /document_categories", getDocuments(serverCfg))
	mux.HandleFunc("GET /document_categories/{id}", getDocumentById(serverCfg))
	mux.HandleFunc("POST /document_categories", auth_helpers.AuthMiddleware(serverCfg, createDocumentCategory(serverCfg)))
	mux.HandleFunc("PUT /document_categories/{id}", auth_helpers.AuthMiddleware(serverCfg, updateDocumentCategory(serverCfg)))
	mux.HandleFunc("DELETE /document_categories/{id}", auth_helpers.AuthMiddleware(serverCfg, deleteDocumentCategory(serverCfg)))
}
