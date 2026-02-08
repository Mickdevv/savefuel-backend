package auth

import "github.com/google/uuid"

type RegisterPayload struct {
	Email     string `json:"email"`
	Password1 string `json:"password_1"`
	Password2 string `json:"password_2"`
}
type RegisterResponse struct {
	Email string    `json:"email"`
	ID    uuid.UUID `json:"id"`
}
type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
type UnsuccessfulLoginResponse struct {
	Error string `json:"error"`
}
type UserWithTokens struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	ID           uuid.UUID `json:"id"`
}
