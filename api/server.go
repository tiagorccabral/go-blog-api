package api

import (
	"github.com/tiagorccabral/go-blog-api/api/controllers"
)

var server = controllers.Server{}

// Run is a wrapper to start routes, DB and the HTTP server itself
func Run() {
	server.Initialize()

	server.RunServer(":8000")
}
