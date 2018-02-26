package routers

import (
	"github.com/kostyaurysov/go-template/controllers"
	"github.com/gorilla/mux"
) 

// SetUserRoutes registers routes for user entity
func SetUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/users", controllers.Register).Methods("POST")
	router.HandleFunc("/users/login", controllers.Login).Methods("POST")
	return router
}
