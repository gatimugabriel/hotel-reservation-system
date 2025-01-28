package middleware

import (
	"github.com/gatimugabriel/hotel-reservation-system/internal/constants"
	"net/http"
	"strings"
)

// RoleCheck checks for required roles before hitting controller
func RoleCheck(allowedRoles []constants.Role, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get user role from context (set by auth middleware)
		userRole := r.Context().Value("role").(string)

		roleInterface := r.Context().Value("role")
		if roleInterface == nil {
			http.Error(w, "Unauthorized - No role found", http.StatusUnauthorized)
			return
		}

		// Check if user role is allowed
		for _, role := range allowedRoles {
			if strings.EqualFold(userRole, string(role)) {
				next.ServeHTTP(w, r)
				return
			}
		}

		// If role is not allowed, return forbidden
		http.Error(w, "Forbidden. You do not have the right permissions to complete the action", http.StatusForbidden)
	})
}

// AuthWithRoleCheck combines Authenticate and RoleCheck middlewares
func AuthWithRoleCheck(allowedRoles []constants.Role) func(http.Handler) http.Handler {
	return ChainMiddleware(
		Authenticate,
		func(next http.Handler) http.Handler {
			return RoleCheck(allowedRoles, next)
		},
	)
}