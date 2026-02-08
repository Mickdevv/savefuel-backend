package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Mickdevv/savefuel-backend/api"
	"github.com/Mickdevv/savefuel-backend/internal/testUtils"
)

func LoginAttemptLimit(t *testing.T, user UserWithTokens, serverCfg *api.ServerConfig) {
	fmt.Println("LoginAttemptLimit")
	for i := range 10 {
		wLogin := httptest.NewRecorder()
		body_login, _ := json.Marshal(LoginPayload{Email: user.Email, Password: "test1234a"})
		req_login := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body_login))
		LoginHandler(serverCfg, wLogin, req_login)

		response := UnsuccessfulLoginResponse{}
		err := json.Unmarshal(wLogin.Body.Bytes(), &response)
		if err != nil {
			t.Fatal(err)
		}
		if wLogin.Code != 401 {
			t.Fatalf("Wrong login code %v at attempt %v", wLogin.Code, i)
		}
	}
	wLogin := httptest.NewRecorder()
	body_login, _ := json.Marshal(LoginPayload{Email: user.Email, Password: user.Password})
	req_login := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body_login))
	LoginHandler(serverCfg, wLogin, req_login)

	response := UnsuccessfulLoginResponse{}
	err := json.Unmarshal(wLogin.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}
	if wLogin.Code != 401 {
		t.Fatalf("Wrong login code %v", wLogin.Code)
	}
}

func TestAuth(t *testing.T) {
	serverCfg := testUtils.TestServerCFG()
	testUser := RegisterAndLogin(t, &serverCfg)
	LoginAttemptLimit(t, testUser, &serverCfg)
	err := CleanupTestUser(testUser.ID, &serverCfg)
	if err != nil {
		t.Fatalf("Failed to clean up test user: %v", err)
	}
}
