package main

import (
	"log"
	"net/http"

	"github.com/softwarearquitecturegroup1a/1A_profile_picture_ms/profilepicture/common"
	"github.com/softwarearquitecturegroup1a/1A_profile_picture_ms/profilepicture/routers"
)

// Entry point for the program
func main() {

	// Calls startup logic
	common.StartUp()
	// Get the mux router object
	router := routers.InitRoutes()

	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: router,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
