package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Server base type. Contains pointer to router
type Server struct {
	Router *mux.Router
}

// Initialize server router and route handlers
func (server *Server) Initialize() {

	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

// RunServer is responsible to start the server
func (server *Server) RunServer(addr string) {
	fmt.Println("Listening to port 8000")

	// Adds logger to server
	loggedRouter := handlers.LoggingHandler(os.Stdout, server.Router)

	log.Fatal(http.ListenAndServe(addr, loggedRouter))
}
