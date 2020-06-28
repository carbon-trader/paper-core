package repository

import (
	"time"

	"github.com/carbon-trader/paper-core/model"
	"gopkg.in/mgo.v2/bson"
)

// PaperService strings to connection with database
type PaperService struct {
	Server   string
	Database string
}

var repository = PaperRepository{}

// Connect with database
func (service *PaperService) Connect() {
	repository._Connect(service.Server, service.Database)
}

//CreateDBIndex create index in database
func (service *PaperService) CreateDBIndex() {
	repository._CreateIndex()
}

//isDBAlive Verify if the database is up
func (service *PaperService) isDBAlive() error {
	return repository._IsAlive()
}

//Save service func
func (service *PaperService) Save(model model.PaperModel) (bson.ObjectId, error) {
	model.ID = bson.NewObjectId()
	model.CreatedAt = time.Now()

	return model.ID, repository._Save(model)
}

//Update service func
func (service *PaperService) Update(id string, model model.PaperModel) error {
	model.UpdatedAt = time.Now()
	return repository._Update(id, model)
}

//Delete service func
func (service *PaperService) Delete(id string) error {
	return repository._Delete(id)
}

//GetAll service func
func (service *PaperService) GetAll() ([]model.PaperModel, error) {
	return repository._GetAll()
}

//FindByPaper service func
func (service *PaperService) FindByPaper(paper string) (model.PaperModel, error) {
	return repository._FindByPaper(paper)
}
