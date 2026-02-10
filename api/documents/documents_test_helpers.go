package documents

import (
	"net/http"
	"testing"

	"github.com/Mickdevv/savefuel-backend/api"
	"github.com/Mickdevv/savefuel-backend/api/auth"
	"github.com/Mickdevv/savefuel-backend/api/document_categories"
)

func UploadDocumentTest(t *testing.T, serverCfg *api.ServerConfig, mux *http.ServeMux, user auth.UserWithTokens, category document_categories.DocumentCategoryResponse) Document {
	return Document{}
}
