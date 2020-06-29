package test

import (
	"testing"
	"time"

	"github.com/carbon-trader/paper-core/model"
	"github.com/carbon-trader/paper-core/repository"
	"gopkg.in/mgo.v2/bson"
)

var service = repository.PaperService{}

const (
	TESTPAPER = "TESTE1"
)

func init() {
	service.Server = "root:example@localhost"
	service.Database = "carbontrader_db"
	service.Connect()
}

func TestAddPaper(t *testing.T) {

	testing.Init()

	var model = model.PaperModel{
		ID:                   bson.NewObjectId(),
		Paper:                "TESTE1",
		Type:                 "ON",
		Company:              "Teste Itau",
		Sector:               "Financeiro",
		Subsector:            "Banco",
		LastBalanceProcessed: "31/03/2020",
		CreatedAt:            time.Now(),
		CreatedBy:            "lribas",
	}

	if _, err := service.Save(model); err != nil {
		t.Errorf("Error to persist information in database %s", err)
	}
}

func TestGetAllPaper(t *testing.T) {
	testing.Init()

	_, err := service.GetAll()

	if err != nil {
		t.Errorf("Error to persist information in database %s", err)
	}
}

func TestFindByPaper(t *testing.T) {
	testing.Init()
	_, err := service.FindByPaper(TESTPAPER)

	if err != nil {
		t.Errorf("Error to persist information in database %s", err)
	}
}

func TestLikeCompany(t *testing.T) {
	testing.Init()

	models, err := service.FindByCompany("Teste")

	if err != nil {
		t.Errorf("Error to execute like query %s", err.Error())
	}

	for _, model := range models {
		if model.Company != "Teste Itau" {
			t.Errorf("Not find elements")
		}
	}
}

func TestFindBySector(t *testing.T) {
	testing.Init()

	_, err := service.FindBySector("Banco")

	if err != nil {
		t.Errorf("Error to find by sector, %s", err.Error())
	}

}

func TestFindBySubsector(t *testing.T) {
	testing.Init()

	_, err := service.FindBySubsector("Financeiro")

	if err != nil {
		t.Errorf("Error to find by sector, %s", err.Error())
	}

}

func TestUpdatePaper(t *testing.T) {
	testing.Init()
	model, err := service.FindByPaper(TESTPAPER)

	if err != nil {
		t.Errorf("Error to persist information in database %s", err)
	} else {

		model.Company = "New Company"

		if err := service.Update(model.ID.Hex(), model); err != nil {
			t.Errorf("Error to update %s", err)
		}
	}
}

func TestDeletePaper(t *testing.T) {
	testing.Init()
	model, err := service.FindByPaper(TESTPAPER)

	if err != nil {
		t.Errorf("Error to persist information in database %s", err)
	} else {

		model.Company = "New Company"

		if err := service.Delete(model.ID.Hex()); err != nil {
			t.Errorf("Error to update %s", err)
		}
	}
}
