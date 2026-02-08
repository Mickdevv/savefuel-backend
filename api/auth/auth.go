package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Mickdevv/savefuel-backend/api"
	"github.com/Mickdevv/savefuel-backend/internal/auth_helpers"
	"github.com/Mickdevv/savefuel-backend/internal/database"
	"github.com/google/uuid"
)

func RegisterHandler(serverCfg *api.ServerConfig, w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email     string `json:"email"`
		Password1 string `json:"password_1"`
		Password2 string `json:"password_2"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		api.RespondWithError(w, http.StatusBadRequest, "Payload error", err)
		return
	}

	if !auth_helpers.ValidateEmail(params.Email) {
		api.RespondWithError(w, http.StatusBadRequest, "Invalid email address", nil)
		return
	}

	if params.Password1 != params.Password2 {
		api.RespondWithError(w, http.StatusBadRequest, "Passwords do not match", nil)
		return
	}

	if err := auth_helpers.ValidatePassword(params.Password1); err != nil {
		api.RespondWithError(w, http.StatusBadRequest, "Invalid password", err)
		return
	}

	hashed_password, err := auth_helpers.HashPassword(params.Password1)
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "Something went wrong", err)
		return
	}
	created_user, err := serverCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		Email:    params.Email,
		Password: hashed_password,
	})
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "Error creating user", err)
		return
	}

	type user struct {
		Email         string    `json:"email"`
		Id            uuid.UUID `json:"id"`
		CreatedAt     time.Time `json:"created_at"`
		UpdatedAt     time.Time `json:"updated_at"`
		EmailVerified bool      `json:"email_verified"`
	}

	u := user{
		Email:         created_user.Email,
		Id:            created_user.ID,
		CreatedAt:     created_user.CreatedAt,
		UpdatedAt:     created_user.UpdatedAt,
		EmailVerified: created_user.EmailVerified,
	}
	api.RespondWithJSON(w, http.StatusOK, u)

}

func LoginHandler(serverCfg *api.ServerConfig, w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		api.RespondWithError(w, http.StatusBadRequest, "Bad request", err)
		return
	}

	user, err := serverCfg.DB.GetUserForAuth(r.Context(), params.Email)
	if err != nil {
		api.RespondWithError(w, http.StatusUnauthorized, "Invalid credentials", err)
		return
	}

	ok, err := auth_helpers.CheckPassword(params.Password, user.Password)
	if err != nil || !ok {
		serverCfg.DB.IncrementLoginAttemptCount(r.Context(), user.ID)
		api.RespondWithError(w, http.StatusUnauthorized, "Invalid credentials", err)
		return
	}

	if user.LoginAttempts >= 10 {
		serverCfg.DB.ResetLoginAttemptCount(r.Context(), user.ID)
		api.RespondWithError(w, http.StatusUnauthorized, "Too many login attempts. Please contact your administrator.", nil)
		return
	}

	token, err := auth_helpers.CreateJWT(serverCfg.JWT_SECRET, user.ID)
	if err != nil {
		serverCfg.DB.IncrementLoginAttemptCount(r.Context(), user.ID)
		api.RespondWithError(w, http.StatusUnauthorized, "Invalid credentials", err)
		return
	}

	refreshToken, err := auth_helpers.CreateRefreshToken(serverCfg, r, user.ID)
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "Internal server error", err)
		return
	}

	type response struct {
		Token        string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	serverCfg.DB.ResetLoginAttemptCount(r.Context(), user.ID)
	api.RespondWithJSON(w, http.StatusOK, response{
		Token:        token,
		RefreshToken: refreshToken,
	})
}

func RefreshTokenHandler(serverCfg *api.ServerConfig, w http.ResponseWriter, r *http.Request) {

	type params struct {
		RefreshToken string `json:"refresh_token"`
	}

	claims, ok := r.Context().Value("claims").(auth_helpers.Claims)
	if !ok {
		http.Error(w, "Unauth_helpers.rized", http.StatusUnauthorized)
		return
	}

	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	parameters := params{}

	err := decoder.Decode(&parameters)
	if err != nil {
		api.RespondWithError(w, http.StatusBadRequest, "Payload error", err)
		return
	}

	oldRefreshToken, err := serverCfg.DB.GetRefreshToken(r.Context(), parameters.RefreshToken)
	if err != nil {
		api.RespondWithError(w, http.StatusNotFound, "Invalid refresh token", err)
		return
	}

	userId, err := uuid.Parse(claims.UserId)
	if err != nil {
		api.RespondWithError(w, http.StatusBadRequest, "Invalid user id", err)
		return
	}

	if userId != oldRefreshToken.UserID {
		api.RespondWithError(w, http.StatusUnauthorized, "Unauthorized", fmt.Errorf("User %v tried to use refresh token for user %v", userId, oldRefreshToken.UserID))
		return
	}

	if oldRefreshToken.RevokedAt.Valid || oldRefreshToken.ExpiresAt.Unix() < time.Now().Unix() {
		api.RespondWithError(w, http.StatusBadRequest, "Invalid refresh token", nil)
		return
	}

	token, err := auth_helpers.CreateJWT(serverCfg.JWT_SECRET, userId)
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "Internal server error", err)
		return
	}

	newRefreshToken, err := auth_helpers.CreateRefreshToken(serverCfg, r, userId)
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "Internal server error", err)
		return
	}

	err = serverCfg.DB.RevokeRefreshToken(r.Context(), oldRefreshToken.Token)
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "Internal server error", err)
		return
	}

	type response struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	api.RespondWithJSON(w, http.StatusOK, response{RefreshToken: newRefreshToken, AccessToken: token})

}
