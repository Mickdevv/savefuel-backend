package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Mickdevv/savefuel-backend/internal/test"
)

type registerPayload struct {
	Email     string `json:"email"`
	Password1 string `json:"password_1"`
	Password2 string `json:"password_2"`
}
type loginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func TestRegisterHandler(t *testing.T) {

	serverCfg := test.TestServerCFG()
	body, _ := json.Marshal(registerPayload{Email: "test@email.com", Password1: "test12345", Password2: "test12345"})
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body))
	w := httptest.NewRecorder()

	registerHandler(&serverCfg, w, req)
}
