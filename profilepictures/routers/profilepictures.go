package routers

import (
	"github.com/gorilla/mux"
	"github.com/softwarearquitecturegroup1a/profilepictures/profilepictures/controllers"
)

func setProfilePictureRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/profilepictures", controllers.GetProfilePictures).Methods("GET")
	router.HandleFunc("/profilepictures", controllers.CreateProfilePicture).Methods("POST")
	router.HandleFunc("/profilepictures/{id}", controllers.GetProfilePictureById).Methods("GET")
	router.HandleFunc("/profilepictures/{id}", controllers.DeleteProfilePicture).Methods("DELETE")
	return router
}
