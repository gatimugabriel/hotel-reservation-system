package utils

import (
	"net/http"
	"strings"
)

// GetResourceIDFromURL retrieves the resource ID from the request
func GetResourceIDFromURL(r *http.Request) string {
	urlParts := strings.Split(r.URL.Path, "/")
	return urlParts[len(urlParts)-1]
}

// GetParamFromURL retrieves a query parameter value from the URL by its name
func GetParamFromURL(r *http.Request, paramName string) string {
	return r.URL.Query().Get(paramName)
}