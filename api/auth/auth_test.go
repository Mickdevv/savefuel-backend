package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Mickdevv/savefuel-backend/internal/auth_helpers"
	"github.com/Mickdevv/savefuel-backend/internal/test"
)

type registerPayload struct {
	Email     string `json:"email"`
	Password1 string `json:"password_1"`
	Password2 string `json:"password_2"`
}
type registerResponse struct {
	Email string `json:"email"`
}
type loginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type loginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func TestRegisterAndlogin(t *testing.T) {

	serverCfg := test.TestServerCFG()
	timestamp := time.Now().Unix()
	user_email := fmt.Sprintf("test-%v@email.com", timestamp)
	user_password := "test12345"

	wRegister := httptest.NewRecorder()

	body_register, _ := json.Marshal(registerPayload{Email: user_email, Password1: user_password, Password2: user_password})
	req_register := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(body_register))
	registerHandler(&serverCfg, wRegister, req_register)

	created_user := registerResponse{}
	err := json.Unmarshal(wRegister.Body.Bytes(), &created_user)
	if err != nil {
		log.Fatal(err)
	}

	wLogin := httptest.NewRecorder()
	body_login, _ := json.Marshal(loginPayload{Email: created_user.Email, Password: user_password})
	req_login := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body_login))
	loginHandler(&serverCfg, wLogin, req_login)

	tokens := loginResponse{}
	err = json.Unmarshal(wLogin.Body.Bytes(), &tokens)
	if err != nil {
		log.Fatal(err)
	}

	_, err = auth_helpers.ValidateJWT(tokens.AccessToken, serverCfg.JWT_SECRET)
	if err != nil {
		log.Fatal(err)
	}

}
