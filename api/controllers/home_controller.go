package controllers

import (
	"net/http"

	"github.com/tiagorccabral/go-blog-api/api/responses"
)

// Home controller is the root route on the API
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome to a REST Go API")
}
