package middleware

import (
	"log"
	"net/http"
	"time"
)

// COLORS
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	//Magenta = "\033[35m"
	//Cyan    = "\033[36m"
)

type responseWriter struct {
	http.ResponseWriter
	StatusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.StatusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// RequestLogger : Console log all http requests
func RequestLogger(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		//	capture status code
		// ww (wrapped writer) / (custom writer)
		ww := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(ww, r)

		//status color
		//statusColor := Green
		//if r.Response.StatusCode >= 400 && r.Response.StatusCode < 500 {
		//	statusColor = Yellow
		//} else if r.Response.StatusCode >= 500 {
		//	statusColor = Red
		//}

		statusColor := Green
		if ww.StatusCode >= 400 && ww.StatusCode < 500 {
			statusColor = Yellow
		} else if ww.StatusCode >= 500 {
			statusColor = Red
		}

		//Log request details
		log.Printf(
			"%s%s%s %s%s%s %s%s%s %s%d%s %s%s%s",
			Blue, r.Method, Reset,
			Yellow, r.RequestURI, Reset,
			Green, r.RemoteAddr, Reset,
			statusColor, ww.StatusCode, Reset,
			Blue, time.Since(start), Reset,
		)
	})
}