package main

import (
	"log"
	"github.com/kostyaurysov/go-template/common"
	"github.com/kostyaurysov/go-template/routers"
	"net/http"
)

func main() {
	// Calls startup logic
	common.StartUp()
	// Get the mux router object
	router := routers.InitRoutes()
	// Create the Server
	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: router,
	}
	log.Println("Listening...")
	// Running the HTTP Server
	server.ListenAndServe()
}
