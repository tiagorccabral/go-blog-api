package controllers

import "github.com/tiagorccabral/go-blog-api/api/middlewares"

func (server *Server) initializeRoutes() {

	server.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(server.Home)).Methods("GET")
}
