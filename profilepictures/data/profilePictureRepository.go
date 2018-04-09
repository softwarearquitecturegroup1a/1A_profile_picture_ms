package data

import (
	"time"
	"github.com/softwarearquitecturegroup1a/profilepictures/profilepictures/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ProfilePictureRepository struct {
	C *mgo.Collection
}

func (r *ProfilePictureRepository) GetAll() []models.ProfilePicture {
	var Profilepictures []models.ProfilePicture
	iter := r.C.Find(nil).Iter()
	result := models.ProfilePicture{}
	for iter.Next(&result) {
		Profilepictures = append(Profilepictures, result)
	}
	return Profilepictures
}

func (r *ProfilePictureRepository) Create(ProfilePicture *models.ProfilePicture) error {
	obj_id := bson.NewObjectId()
	ProfilePicture.Id = obj_id
	ProfilePicture.CreatedOn = time.Now()
	err := r.C.Insert(&ProfilePicture)
	return err
}

func (r *ProfilePictureRepository) GetById(id string) (ProfilePicture models.ProfilePicture, err error) {
	err = r.C.FindId(bson.ObjectIdHex(id)).One(&ProfilePicture)
	return
}

func (r *ProfilePictureRepository) GetByIdStudent(id string) (ProfilePicture models.ProfilePicture, err error) {
	err = r.C.Find(bson.M{"idStudent": id} ).One(&ProfilePicture)
	return
}

func (r *ProfilePictureRepository) Delete(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
