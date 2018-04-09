package controllers

import (
	"github.com/softwarearquitecturegroup1a/profilepictures/profilepictures/models"
)

type (
	// For Get - /profilepicture
	ProfilePicturesResource struct {
		Data []models.ProfilePicture `json:"data"`
	}
	// For Post/Put - /profilepicture
	ProfilePictureResource struct {
		Data models.ProfilePicture `json:"data"`
	}
)
