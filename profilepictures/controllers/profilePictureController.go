package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/softwarearquitecturegroup1a/profilepictures/profilepictures/common"
	"github.com/softwarearquitecturegroup1a/profilepictures/profilepictures/data"
	"gopkg.in/mgo.v2"
)

// Handler for HTTP Get - "/ProfilePictures"
// Returns all ProfilePicture documents
func GetProfilePictures(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("ProfilePictures")
	repo := &data.ProfilePictureRepository{c}
	ProfilePictures := repo.GetAll()
	j, err := json.Marshal(ProfilePicturesResource{Data: ProfilePictures})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Handler for HTTP Post - "/ProfilePictures"
// Insert a new ProfilePicture document
func CreateProfilePicture(w http.ResponseWriter, r *http.Request) {
	var dataResourse ProfilePictureResource
	// Decode the incoming ProfilePicture json
	err := json.NewDecoder(r.Body).Decode(&dataResourse)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid ProfilePicture data", 500)
		return
	}
	ProfilePicture := &dataResourse.Data

	// create new context
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("ProfilePictures")
	// Insert a ProfilePicture document
	repo := &data.ProfilePictureRepository{c}
	repo.Create(ProfilePicture)
	j, err := json.Marshal(dataResourse)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Handler for HTTP Get - "/ProfilePictures/{id}"
// Get ProfilePicture by id
func GetProfilePictureById(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// create new context
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("ProfilePictures")
	repo := &data.ProfilePictureRepository{c}

	// Get ProfilePicture by id
	ProfilePicture, err := repo.GetByIdStudent(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		} else {
			common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
			return
		}
	}

	j, err := json.Marshal(ProfilePicture)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Handler for HTTP Delete - "/ProfilePictures/{id}"
// Delete ProfilePicture by id
func DeleteProfilePicture(rw http.ResponseWriter, req *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(req)
	id := vars["id"]

	// Create new context
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("ProfilePictures")
	repo := &data.ProfilePictureRepository{c}

	err := repo.Delete(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			rw.WriteHeader(http.StatusNotFound)
			return
		} else {
			common.DisplayAppError(rw, err, "An unexpected error has occurred", 500)
			return
		}
	}
	rw.WriteHeader(http.StatusNoContent)
}
