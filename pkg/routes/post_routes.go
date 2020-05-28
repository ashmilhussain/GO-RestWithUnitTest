package routes

import "github.com/ashmilhussain/GO-RestWithUnitTest/pkg/middlewares"

func (server *Server) addPostRoutes() {
	server.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(server.Handler.CreatePost)).Methods("POST")
	server.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(server.Handler.GetPosts)).Methods("GET")
	server.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(server.Handler.UpdatePost, server.Handler.DB))).Methods("PUT")
	server.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareAuthentication(server.Handler.DeletePost, server.Handler.DB)).Methods("DELETE")
}
