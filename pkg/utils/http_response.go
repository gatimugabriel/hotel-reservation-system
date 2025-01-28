package utils

import (
	"encoding/json"
	"net/http"
)

// RespondJSON : helper function to write JSON responses
func RespondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// RespondError : helper function to write error responses
func RespondError(w http.ResponseWriter, code int, message string) {
	RespondJSON(w, code, map[string]string{"error": message})
}

type PaginatedResponseData struct {
	Data  interface{} `json:"data"`
	Count int         `json:"count"`
	Page  int         `json:"page"`
	Total int         `json:"total"`
}

func RespondPaginatedJSON(w http.ResponseWriter, status int, data interface{}, count int, page int, total int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(PaginatedResponseData{
		Data:  data,
		Count: count,
		Page:  page,
		Total: total,
	})
	if err != nil {
		return
	}
}