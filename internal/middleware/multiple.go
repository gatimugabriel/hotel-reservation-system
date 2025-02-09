package middleware

import "net/http"

// ApplyMiddlewareToMany applies a middleware to multiple handlers
// @returns: a slice of wrapped handlers
func ApplyMiddlewareToMany(middleware func(http.Handler) http.Handler, handlers ...http.HandlerFunc) []http.Handler {
	wrappedHandlers := make([]http.Handler, len(handlers))
	for i, handler := range handlers {
		wrappedHandlers[i] = middleware(handler)
	}
	return wrappedHandlers
}