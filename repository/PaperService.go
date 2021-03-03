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

//IsDBAlive Verify if the database is up
func (service *PaperService) IsDBAlive() error {
	return repository._IsAlive()
}

//Save service func
func (service *PaperService) Save(model model.Paper) (bson.ObjectId, error) {
	model.ID = bson.NewObjectId()
	model.CreatedAt = time.Now()

	return model.ID, repository._Save(model)
}

//Update service func
func (service *PaperService) Update(id string, model model.Paper) error {
	model.UpdatedAt = time.Now()
	return repository._Update(id, model)
}

//Delete service func
func (service *PaperService) Delete(id string) error {
	return repository._Delete(id)
}

//GetAll service func
func (service *PaperService) GetAll() ([]model.Paper, error) {
	return repository._GetAll()
}

//FindByPaper service func
func (service *PaperService) FindByPaper(paper string) (model.Paper, error) {
	return repository._FindByPaper(paper)
}

//FindByCommercialName Apply Like %TEXT%
func (service *PaperService) FindByCommercialName(commercialName string) ([]model.Paper, error) {
	return repository._FindByCommercialName(commercialName)
}

//FindBySector find paper by sector
func (service *PaperService) FindBySector(sector string) ([]model.Paper, error) {
	return repository._FindBySector(sector)
}

//FindBySubsector find paper by subsector
func (service *PaperService) FindBySubsector(subsector string) ([]model.Paper, error) {
	return repository._FindBySubSector(subsector)
}

//FindByID find paper by ID
func (service *PaperService) FindByID(id string) (model.Paper, error) {
	return repository._FindByID(id)
}
