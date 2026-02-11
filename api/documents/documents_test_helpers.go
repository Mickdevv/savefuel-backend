package documents

import (
	"bytes"
	"encoding/json"
	"fmt"
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

	// filePath := path.Join(serverCfg.STATIC_FILES_DIR, "documents", "installed-packages.txt")
	filePath := path.Join("file1.txt")
	err := os.WriteFile(filePath, []byte("test"), 0644)
	file, err := os.Open(filePath)
	if err != nil {
		t.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	body := bytes.Buffer{}
	writer := multipart.NewWriter(&body)

	metadata, _ := json.Marshal(document)

	writer.WriteField("metadata", string(metadata))
	fmt.Println(string(metadata))

	part, err := writer.CreateFormFile("file", filePath)
	if err != nil {
		t.Fatalf("Error creating form file from file: %v", err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		t.Fatalf("Error copying file to body: %v", err)
	}
	writer.Close()

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/documents", &body)
	r.Header.Set("Content-type", writer.FormDataContentType())
	r.Header.Set("Authorization", "Bearer "+user.AccessToken)

	mux.ServeHTTP(w, r)
	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d, body: %s", w.Code, w.Body.String())
	}

	fmt.Println(w.Body)
	doc := UploadDocumentResponse{}
	decoder := json.NewDecoder(w.Body)
	err = decoder.Decode(&doc)
	if err != nil {
		t.Fatalf("Error decoding response: %v", err)
	}

	return doc.Data
}
