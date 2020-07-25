package middlewares

import (
	"net/http"
)

// SetMiddlewareJSON is responsible to set application content-type on JSON responses on the API
func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}
