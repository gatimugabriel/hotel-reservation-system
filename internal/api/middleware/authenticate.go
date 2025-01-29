package middleware

import (
	"context"
	"github.com/gatimugabriel/hotel-reservation-system/pkg/utils"
	"log"
	"net/http"
	"strings"
)

// Authenticate -> extracts userID from access token & adds it to context
func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			if cookie, err := r.Cookie("accessToken"); err == nil {
				authHeader = "Bearer " + cookie.Value
			}
		}

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			utils.RespondError(w, http.StatusUnauthorized, "Missing Access Token")
			return
		}

		accessToken := strings.TrimPrefix(authHeader, "Bearer ")

		//	 Validate token
		payload, err := utils.ValidateToken(accessToken, "ACCESS")
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Invalid Token")
			return
		}

		userID := payload["userID"].(string)
		role := payload["role"].(string)

		// Set user values to content
		ctx := context.WithValue(r.Context(), "userID", userID)
		ctx = context.WithValue(ctx, "role", role)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}