package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/tiagorccabral/go-blog-api/api/models"

	//postgres database driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Server base type. Contains pointer to router
type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

// Initialize server router and route handlers.
// Params (string): Driver, User, Password, Port, Host, DBName
func (server *Server) Initialize(DBDriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error

	if DBDriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		server.DB, err = gorm.Open(DBDriver, DBURL)
		if err != nil {
			fmt.Printf("Failed to connect to %s database\n", DBDriver)
			log.Fatal("The following error happened: ", err)
		} else {
			log.Printf("Connected to %s database. DB name: %s\n", DBDriver, DbName)
		}
	}

	// Generates DB mirations
	server.DB.Debug().AutoMigrate(&models.User{}, &models.Post{})

	// Initializes Router
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
