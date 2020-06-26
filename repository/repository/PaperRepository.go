package repository

import (
	"log"

	"github.com/carbon-trader/paper-core/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// PaperRepository struct
type PaperRepository struct {
	Server   string
	Database string
}

var db *mgo.Database

// const
const (
	COLLECTION = "paper"
)

//Connect function to access database
func (p *PaperRepository) Connect() {
	session, err := mgo.Dial(p.Server)

	// Verify if occurs some error and stop application
	if err != nil {
		log.Fatal(err)
	}

	// Establish a session with database
	db = session.DB(p.Database)
}

// Save func
func (p *PaperRepository) Save(model model.PaperModel) error {
	err := db.C(COLLECTION).Insert(&model)
	return err
}

// Delete func
func (p *PaperRepository) Delete(id string) error {
	err := db.C(COLLECTION).RemoveId(bson.ObjectIdHex(id))
	return err
}

// Update func
func (p *PaperRepository) Update(id string, model model.PaperModel) error {
	err := db.C(COLLECTION).UpdateId(bson.ObjectIdHex(id), &model)
	return err
}

// GetAll func
func (p *PaperRepository) GetAll() ([]model.PaperModel, error) {
	var model []model.PaperModel
	err := db.C(COLLECTION).Find(bson.M{}).All(&model)
	return model, err
}
