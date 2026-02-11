package documents

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Mickdevv/savefuel-backend/api"
	"github.com/Mickdevv/savefuel-backend/api/auth"
	"github.com/Mickdevv/savefuel-backend/api/document_categories"
	"github.com/Mickdevv/savefuel-backend/internal/testUtils"
)

func UpdateDocumentTest(t *testing.T, serverCfg *api.ServerConfig, mux *http.ServeMux, user auth.UserWithTokens, document Document) Document {
	w := httptest.NewRecorder()
	body, _ := json.Marshal(DocumentPayload{
		Locale:      "EN",
		Title:       "Test document",
		Description: "This is a test document",
		Priority:    10,
		Active:      false,
		CategoryID:  document.CategoryID,
	})
	r := httptest.NewRequest(http.MethodPost, "/documents", bytes.NewReader(body))
	r.Header.Add("Authorization", "Bearer "+user.AccessToken)

	mux.ServeHTTP(w, r)

	doc := Document{}
	err := json.Unmarshal(w.Body.Bytes(), &doc)

	if err != nil {
		t.Fatalf("Error creating document: %v", err)
	}
	if w.Code != 200 {
		t.Fatalf("Error creating document: %v", w.Body)
	}
	return doc

}

func TestDocuments(t *testing.T) {
	serverCfg := testUtils.TestServerCFG()
	mux := http.NewServeMux()
	RegisterRoutes(mux, &serverCfg)
	auth.RegisterRoutes(mux, &serverCfg)
	document_categories.RegisterRoutes(mux, &serverCfg)

	user := auth.RegisterAndLogin(t, &serverCfg, mux)
	category := document_categories.CreateDocumentCategoryTest(t, serverCfg, mux, user)

	document := UploadDocumentTest(t, &serverCfg, mux, user, category)
	fmt.Println("Document uploaded successfully")
	fmt.Println(document)
	// document = UpdateDocumentTest(t, &serverCfg, mux, user, document)
	// fmt.Println(document)

	auth.CleanupTestUser(user.ID, &serverCfg)
}
