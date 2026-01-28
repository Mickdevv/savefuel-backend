package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/Mickdevv/savefuel-backend/api"
)

func AuthMiddleware(serverCfg *api.ServerConfig, next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Auth middleware")

		authorizationHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix("Bearer ", authorizationHeader) {
			api.RespondWithError(w, http.StatusUnauthorized, "Token header malformed", nil)
			return
		}

		tokenString := strings.TrimPrefix("Bearer ", authorizationHeader)

		claims, err := ValidateJWT(tokenString, serverCfg.JWT_SECRET)
		if err != nil {
			api.RespondWithError(w, http.StatusUnauthorized, "Unauthorized", err)
			return
		}

		fmt.Println("Auth middleware", claims)
		ctx := context.WithValue(r.Context(), "claims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
