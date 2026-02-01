package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Mickdevv/savefuel-backend/api"
	"github.com/Mickdevv/savefuel-backend/internal/auth"
	"github.com/Mickdevv/savefuel-backend/internal/database"
	"github.com/google/uuid"
)

func registerHandler(serverCfg *api.ServerConfig, w http.ResponseWriter, r *http.Request) {
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

	if !auth.ValidateEmail(params.Email) {
		api.RespondWithError(w, http.StatusBadRequest, "Invalid email address", nil)
		return
	}

	if params.Password1 != params.Password2 {
		api.RespondWithError(w, http.StatusBadRequest, "Passwords do not match", nil)
		return
	}

	if err := auth.ValidatePassword(params.Password1); err != nil {
		api.RespondWithError(w, http.StatusBadRequest, "Invalid password", err)
		return
	}

	hashed_password, err := auth.HashPassword(params.Password1)
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

func loginHandler(serverCfg *api.ServerConfig, w http.ResponseWriter, r *http.Request) {
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

	ok, err := auth.CheckPassword(params.Password, user.Password)
	if err != nil || !ok {
		api.RespondWithError(w, http.StatusUnauthorized, "Invalid credentials", err)
		return
	}

	token, err := auth.CreateJWT(serverCfg.JWT_SECRET, user.ID)
	if err != nil {
		api.RespondWithError(w, http.StatusUnauthorized, "Invalid credentials", err)
		return
	}

	refreshToken, err := auth.CreateRefreshToken(serverCfg, r, user.ID)
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "Internal server error", err)
		return
	}

	type response struct {
		Token        string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	api.RespondWithJSON(w, http.StatusOK, response{
		Token:        token,
		RefreshToken: refreshToken,
	})

}

func refreshTokenHandler(serverCfg *api.ServerConfig, w http.ResponseWriter, r *http.Request) {

	// Verify user
	// Verify refresh token validity
	// Generate new tokens
	// Revoke old tokens
	// Return new tokens

	type params struct {
		RefreshToken string `json:"refresh_token"`
	}

	claims, ok := r.Context().Value("claims").(auth.Claims)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
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

	token, err := auth.CreateJWT(serverCfg.JWT_SECRET, userId)
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "Internal server error", err)
		return
	}

	newRefreshToken, err := auth.CreateRefreshToken(serverCfg, r, userId)
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
