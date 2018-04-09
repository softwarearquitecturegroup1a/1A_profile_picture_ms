package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	ProfilePicture struct {
		Id        bson.ObjectId `bson:"_id,omitempty" json:"id"`
		IdStudent string        `bson:"idStudent" json:"Student"`
		UrlPhoyo  string        `json:"Url"`
		CreatedOn time.Time     `json:"createdon,omitempty"`
	}
)
