package middleware

import (
	"log/slog"
	"net/http"
)

func ErrorLogger(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rr := &responseRecorder{ResponseWriter: w, statusCode: http.StatusOK}
			next.ServeHTTP(rr, r)
			if rr.statusCode >= 400 {
				logger.Error("HTTP error", "status", rr.statusCode, "method", r.Method, "url", r.URL.String())
			}
		})
	}
}

type responseRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (rr *responseRecorder) WriteHeader(statusCode int) {
	rr.statusCode = statusCode
	rr.ResponseWriter.WriteHeader(statusCode)
}
