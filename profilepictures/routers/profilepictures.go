package routers

import (
	"github.com/gorilla/mux"
	"github.com/softwarearquitecturegroup1a/1A_profile_picture_ms/profilepictures/controllers"
)

func setProfilepictureRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/Profilepictures", controllers.GetProfilepictures).Methods("GET")
	router.HandleFunc("/Profilepictures", controllers.CreateProfilepicture).Methods("POST")
	router.HandleFunc("/Profilepictures/{id}", controllers.GetProfilepictureById).Methods("GET")
	router.HandleFunc("/Profilepictures/{id}", controllers.DeleteProfilepicture).Methods("DELETE")
	return router
}
