package routes

import (
	"fmt"
	"log"
	"net/http"

	myHandler "github.com/ashmilhussain/GO-RestWithUnitTest/pkg/handlers"
	"github.com/ashmilhussain/GO-RestWithUnitTest/pkg/middlewares"
	"github.com/gorilla/mux"
)

type Server struct {
	Handler myHandler.Server
	Router  *mux.Router
}

func (server *Server) InitializeRoutes() {

	server.Router = mux.NewRouter()

	// Login Route
	server.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(server.Handler.Login)).Methods("POST")

	//Users routes
	server.addUserRoutes()

	//Topic routes
	server.addPostRoutes()

	// Home Route
	server.Router.PathPrefix("/").Handler(http.FileServer(http.Dir("./web/")))

}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
