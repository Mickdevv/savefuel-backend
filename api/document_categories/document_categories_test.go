package document_categories

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Mickdevv/savefuel-backend/api"
	"github.com/Mickdevv/savefuel-backend/api/auth"
	"github.com/Mickdevv/savefuel-backend/internal/testUtils"
	"github.com/google/uuid"
)

func GetDocumentCategoriesTest(t *testing.T, serverCfg api.ServerConfig, mux *http.ServeMux, user auth.UserWithTokens) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/document-categories", nil)

	mux.ServeHTTP(w, r)

	response := []DocumentCategoryResponse{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Response body error: %v", err)
	}
	if w.Code != 200 {
		t.Fatalf("Response error: %v", w.Body)
	}
}

func GetDocumentCategoryByIdTest(t *testing.T, serverCfg api.ServerConfig, mux *http.ServeMux, user auth.UserWithTokens, id uuid.UUID) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/document-categories/%v", id.String()), nil)

	mux.ServeHTTP(w, r)

	response := DocumentCategoryResponse{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Response body error: %v", err)
	}
	if w.Code != 200 {
		t.Fatalf("Response error: %v", w.Body)
	}
}

func UpdateDocumentCategoryTest(t *testing.T, serverCfg api.ServerConfig, mux *http.ServeMux, user auth.UserWithTokens, documentCategory DocumentCategoryResponse) DocumentCategoryResponse {
	w := httptest.NewRecorder()
	body, _ := json.Marshal(UpdateDocumentCategoryPayload{Name: "Test category updated", Active: false})
	r := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/document-categories/%v", documentCategory.ID.String()), bytes.NewReader(body))
	r.Header.Add("Authorization", "Bearer "+user.AccessToken)

	mux.ServeHTTP(w, r)

	response := DocumentCategoryResponse{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}
	if w.Code != 200 {
		t.Fatalf("Response error: %v", w.Body)
	}

	return response
}

func DeleteDocumentCategoryTest(t *testing.T, serverCfg api.ServerConfig, mux *http.ServeMux, user auth.UserWithTokens, id uuid.UUID) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/document-categories/%v", id.String()), nil)
	r.Header.Add("Authorization", "Bearer "+user.AccessToken)

	mux.ServeHTTP(w, r)
	if w.Code != 200 {
		t.Fatalf("Response error: %v", w.Body)
	}
}

func TestDocumentCategories(t *testing.T) {
	serverCfg := testUtils.TestServerCFG()
	mux := http.NewServeMux()
	RegisterRoutes(mux, &serverCfg)
	auth.RegisterRoutes(mux, &serverCfg)

	user := auth.RegisterAndLogin(t, &serverCfg, mux)

	documentCategory := CreateDocumentCategoryTest(t, serverCfg, mux, user)
	documentCategory = UpdateDocumentCategoryTest(t, serverCfg, mux, user, documentCategory)
	GetDocumentCategoriesTest(t, serverCfg, mux, user)
	GetDocumentCategoryByIdTest(t, serverCfg, mux, user, documentCategory.ID)
	DeleteDocumentCategoryTest(t, serverCfg, mux, user, documentCategory.ID)

	auth.CleanupTestUser(user.ID, &serverCfg)
}
