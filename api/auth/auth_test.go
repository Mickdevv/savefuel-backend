package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Mickdevv/savefuel-backend/api"
	"github.com/Mickdevv/savefuel-backend/internal/testUtils"
)

func LoginAttemptLimit(t *testing.T, mux *http.ServeMux, user UserWithTokens, serverCfg *api.ServerConfig) {
	for i := range 10 {
		w := httptest.NewRecorder()
		body, _ := json.Marshal(LoginPayload{Email: user.Email, Password: "test1234a"})
		r := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body))

		mux.ServeHTTP(w, r)

		response := UnsuccessfulLoginResponse{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		if err != nil {
			t.Fatal(err)
		}
		if w.Code != 401 {
			t.Fatalf("Wrong login code %v at attempt %v", w.Code, i)
		}
	}
	w := httptest.NewRecorder()
	body, _ := json.Marshal(LoginPayload{Email: user.Email, Password: user.Password})
	r := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body))
	LoginHandler(serverCfg, w, r)

	response := UnsuccessfulLoginResponse{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}
	if w.Code != 401 {
		t.Fatalf("Wrong login code %v", w.Code)
	}
}

func TestAuth(t *testing.T) {
	serverCfg := testUtils.TestServerCFG()
	mux := http.NewServeMux()
	RegisterRoutes(mux, &serverCfg)

	testUser := RegisterAndLogin(t, &serverCfg, mux)
	LoginAttemptLimit(t, mux, testUser, &serverCfg)

	err := CleanupTestUser(testUser.ID, &serverCfg)
	if err != nil {
		t.Fatalf("Failed to clean up test user: %v", err)
	}
}
