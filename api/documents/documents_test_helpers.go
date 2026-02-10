package documents

import (
	"net/http"
	"net/http/httptest"
	"path"
	"testing"

	"github.com/Mickdevv/savefuel-backend/api"
	"github.com/Mickdevv/savefuel-backend/api/auth"
	"github.com/Mickdevv/savefuel-backend/api/document_categories"
)

func UploadDocumentTest(t *testing.T, serverCfg *api.ServerConfig, mux *http.ServeMux, user auth.UserWithTokens, category document_categories.DocumentCategoryResponse) Document {
	document := UploadDocumentPayload{

		Title:       "Test document title",
		Locale:      "EN",
		Description: "This is a test document description ",
		Priority:    0,
		CategoryID:  category.ID,
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/documents", nil)
	r.FormFile(path.Join(serverCfg.STATIC_FILES_DIR, "documents", "installed-packages.txt"))
	r.MultipartForm.File()

	r.MultipartForm.Value
	return Document{}
}
