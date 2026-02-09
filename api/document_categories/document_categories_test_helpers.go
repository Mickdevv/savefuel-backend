package document_categories

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Mickdevv/savefuel-backend/api"
	"github.com/Mickdevv/savefuel-backend/api/auth"
)

func CreateDocumentCategoryTest(t *testing.T, serverCfg api.ServerConfig, mux *http.ServeMux, user auth.UserWithTokens) DocumentCategoryResponse {
	w := httptest.NewRecorder()
	body, _ := json.Marshal(CreateDocumentCategoryPayload{Name: "Test category"})
	r := httptest.NewRequest(http.MethodPost, "/document-categories", bytes.NewReader(body))
	r.Header.Add("Authorization", "Bearer "+user.AccessToken)

	mux.ServeHTTP(w, r)

	response := DocumentCategoryResponse{}
	if w.Code != 200 {
		t.Fatalf("Response error: %v", w.Body)
	}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	return response
}
