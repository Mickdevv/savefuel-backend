package documents

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
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

	filePath := path.Join(serverCfg.STATIC_FILES_DIR, "documents", "installed-packages.txt")
	file, err := os.Open(filePath)
	if err != nil {
		t.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writer.WriteField("title", document.Title)
	writer.WriteField("locale", document.Locale)
	writer.WriteField("description", document.Description)
	writer.WriteField("priority", document.Priority)
	writer.WriteField("category_id", document.CategoryID.String())

	part, err := writer.CreateFormFile("file", "installed-packages.txt")
	if err != nil {
		t.Fatalf("Error creating form file from file: %v", err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		t.Fatalf("honestly not sure what this part does but whatever", err)
	}

	writer.Close()

	return Document{}
}
