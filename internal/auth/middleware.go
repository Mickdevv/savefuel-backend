package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/Mickdevv/savefuel-backend/api"
)

func AuthMiddleware(serverCfg *api.ServerConfig, next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		authorizationHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authorizationHeader, "Bearer ") {
			api.RespondWithError(w, http.StatusUnauthorized, "Token header malformed", nil)
			return
		}
		tokenString := strings.TrimPrefix(authorizationHeader, "Bearer ")

		claims, err := ValidateJWT(tokenString, serverCfg.JWT_SECRET)
		if err != nil {
			api.RespondWithError(w, http.StatusUnauthorized, "Unauthorized", err)
			return
		}

		ctx := context.WithValue(r.Context(), "claims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
