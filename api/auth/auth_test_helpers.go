package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Mickdevv/savefuel-backend/api"
	"github.com/Mickdevv/savefuel-backend/internal/auth_helpers"
	"github.com/google/uuid"
)

func RegisterAndLogin(t *testing.T, serverCfg *api.ServerConfig) UserWithTokens {

	timestamp := time.Now().Unix()
	user_email := fmt.Sprintf("test-%v@email.com", timestamp)
	user_password := "test12345"

	wRegister := httptest.NewRecorder()
	body_register, _ := json.Marshal(RegisterPayload{Email: user_email, Password1: user_password, Password2: user_password})
	req_register := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(body_register))
	RegisterHandler(serverCfg, wRegister, req_register)

	created_user := RegisterResponse{}
	err := json.Unmarshal(wRegister.Body.Bytes(), &created_user)
	if err != nil {
		t.Fatal(err)
	}

	wLogin := httptest.NewRecorder()
	body_login, _ := json.Marshal(LoginPayload{Email: created_user.Email, Password: user_password})
	req_login := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body_login))
	LoginHandler(serverCfg, wLogin, req_login)

	tokens := LoginResponse{}
	err = json.Unmarshal(wLogin.Body.Bytes(), &tokens)
	if err != nil {
		t.Fatal(err)
	}

	_, err = auth_helpers.ValidateJWT(tokens.AccessToken, serverCfg.JWT_SECRET)
	if err != nil {
		t.Fatal(err)
	}
	return UserWithTokens{ID: created_user.ID, Email: created_user.Email, Password: user_password, RefreshToken: tokens.RefreshToken, AccessToken: tokens.AccessToken}
}

func CleanupTestUser(userId uuid.UUID, serverCfg *api.ServerConfig) error {
	return serverCfg.DB.Deleteuser(context.Background(), userId)
}
