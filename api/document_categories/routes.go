package document_categories

import (
	"net/http"

	"github.com/Mickdevv/savefuel-backend/api"
	"github.com/Mickdevv/savefuel-backend/internal/auth_helpers"
)

func RegisterRoutes(mux *http.ServeMux, serverCfg *api.ServerConfig) {
	mux.HandleFunc("GET /document-categories", GetDocumentCategories(serverCfg))
	mux.HandleFunc("GET /document-categories/{id}", GetDocumentCategoryById(serverCfg))
	mux.HandleFunc("POST /document-categories", auth_helpers.AuthMiddleware(serverCfg, CreateDocumentCategory(serverCfg)))
	mux.HandleFunc("PUT /document-categories/{id}", auth_helpers.AuthMiddleware(serverCfg, UpdateDocumentCategory(serverCfg)))
	mux.HandleFunc("DELETE /document-categories/{id}", auth_helpers.AuthMiddleware(serverCfg, DeleteDocumentCategory(serverCfg)))
}
