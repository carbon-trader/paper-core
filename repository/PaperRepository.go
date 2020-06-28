package repository

import (
	"log"

	"github.com/carbon-trader/paper-core/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// PaperRepository exporter
type PaperRepository struct{}

var db *mgo.Database

// const
const (
	COLLECTION = "paper"
)

func (repository *PaperRepository) _IsAlive() error {
	return db.Session.Ping()
}

func (repository *PaperRepository) _CreateIndex() {
	index := mgo.Index{
		Key:        []string{"paper", "company"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	if err := db.C(COLLECTION).EnsureIndex(index); err != nil {
		//stop the execution
		panic(err)
	}
}

func (repository *PaperRepository) _Connect(Server string, Database string) {
	session, err := mgo.Dial(Server)

	// Verify if occurs some error and stop application
	if err != nil {
		log.Fatal(err)
	}

	// Establish a session with database
	db = session.DB(Database)
}

func (repository *PaperRepository) _Save(model model.PaperModel) error {
	err := db.C(COLLECTION).Insert(&model)
	return err
}

func (repository *PaperRepository) _Update(id string, model model.PaperModel) error {
	err := db.C(COLLECTION).UpdateId(bson.ObjectIdHex(id), &model)
	return err
}

func (repository *PaperRepository) _Delete(id string) error {
	err := db.C(COLLECTION).RemoveId(bson.ObjectIdHex(id))
	return err
}

func (repository *PaperRepository) _GetAll() ([]model.PaperModel, error) {
	var model []model.PaperModel
	err := db.C(COLLECTION).Find(bson.M{}).All(&model)

	return model, err
}

func (repository *PaperRepository) _FindByPaper(paper string) (model.PaperModel, error) {
	var model model.PaperModel
	err := db.C(COLLECTION).Find(bson.M{"paper": paper}).One(&model)
	return model, err
}
