package data

import (
	"time"

	"github.com/softwarearquitecturegroup1a/1A_profile_picture_ms/profilepicture/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ProfilepictureRepository struct {
	C *mgo.Collection
}

func (r *ProfilepictureRepository) GetAll() []models.Profilepicture {
	var Profilepictures []models.Profilepicture
	iter := r.C.Find(nil).Iter()
	result := models.Profilepicture{}
	for iter.Next(&result) {
		Profilepictures = append(Profilepictures, result)
	}
	return Profilepictures
}

func (r *ProfilepictureRepository) Create(Profilepicture *models.Profilepicture) error {
	obj_id := bson.NewObjectId()
	Profilepicture.Id = obj_id
	Profilepicture.CreatedOn = time.Now()
	err := r.C.Insert(&Profilepicture)
	return err
}

func (r *ProfilepictureRepository) GetById(id string) (Profilepicture models.Profilepicture, err error) {
	err = r.C.FindId(bson.ObjectIdHex(id)).One(&Profilepicture)
	return
}

func (r *ProfilepictureRepository) Delete(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
