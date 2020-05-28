package routes

import "github.com/ashmilhussain/GO-RestWithUnitTest/pkg/middlewares"

func (server *Server) addUserRoutes() {
	server.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(server.Handler.CreateUser)).Methods("POST")
	server.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(server.Handler.GetUsers, server.Handler.DB))).Methods("GET")
	server.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(server.Handler.GetUser, server.Handler.DB))).Methods("GET")
	server.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(server.Handler.UpdateUser, server.Handler.DB))).Methods("PUT")
	server.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(server.Handler.DeleteUser, server.Handler.DB)).Methods("DELETE")
}
