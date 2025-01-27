package middleware

import (
	"net/http"
	"strings"
)

func RoleCheck(allowedRoles []string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get user role from context (set by auth middleware)
		userRole := r.Context().Value("userRole").(string)
		
		// Check if user role is allowed
		for _, role := range allowedRoles {
			if strings.EqualFold(userRole, role) {
				next.ServeHTTP(w, r)
				return
			}
		}
		
		// If role is not allowed, return forbidden
		http.Error(w, "Forbidden", http.StatusForbidden)
	})
}