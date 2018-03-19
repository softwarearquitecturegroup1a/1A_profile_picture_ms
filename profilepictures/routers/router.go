package routers

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	// Routes for the profile picture entity
	router = setProfilePictureRouters(router)
	return router
}
