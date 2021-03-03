package test

import (
	"testing"
	"time"

	"github.com/carbon-trader/paper-core/model"
	"github.com/carbon-trader/paper-core/repository"
	"gopkg.in/mgo.v2/bson"
)

var service = repository.PaperService{}

func init() {
	service.Server = "root:root@localhost"
	service.Database = "carbontrader_db"
	service.Connect()
}

func TestAddPaper(t *testing.T) {

	testing.Init()

	var model = model.Paper{
		ID:             bson.NewObjectId(),
		Paper:          "ITSA4",
		CommercialName: "ITAÚSA",
		CompanyName:    "ITAÚSA - INVESTIMENTOS ITAÚ S.A.",
		Type:           "PN N1",
		Sector:         "Financeiros",
		SubSector:      "Bancos",
		MarketValue:    "96.303.800.000",
		CompanyValue:   "97.143.800.000",
		NStock:         "8.410.810.000",
		LastBalance:    time.Now(),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
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
	_, err := service.FindByPaper("ITSA4")

	if err != nil {
		t.Errorf("Error to persist information in database %s", err)
	}
}

func TestLikeCompany(t *testing.T) {
	testing.Init()

	models, err := service.FindByCommercialName("ITAÚSA")

	if err != nil {
		t.Errorf("Error to execute like query %s", err.Error())
	}

	for _, model := range models {
		if model.CommercialName != "ITAÚSA" {
			t.Errorf("Not find elements")
		}
	}
}

func TestFindBySector(t *testing.T) {
	testing.Init()

	_, err := service.FindBySector("Financeiros")

	if err != nil {
		t.Errorf("Error to find by sector, %s", err.Error())
	}

}

func TestFindBySubsector(t *testing.T) {
	testing.Init()

	_, err := service.FindBySubsector("Bancos")

	if err != nil {
		t.Errorf("Error to find by sector, %s", err.Error())
	}

}

func TestUpdatePaper(t *testing.T) {
	testing.Init()
	model, err := service.FindByPaper("ITSA4")

	if err != nil {
		t.Errorf("Error to persist information in database %s", err)
	} else {

		model.CommercialName = "ITAÚSA UPDATED"

		if err := service.Update(model.ID.Hex(), model); err != nil {
			t.Errorf("Error to update %s", err)
		}
	}
}

func TestDeletePaper(t *testing.T) {
	testing.Init()
	model, err := service.FindByPaper("ITSA4")

	if err != nil {
		t.Errorf("Error to persist information in database %s", err)
	} else {

		if model.CommercialName == "ITAÚSA UPDATED" {
			if err := service.Delete(model.ID.Hex()); err != nil {
				t.Errorf("Error to update %s", err)
			}
		}
	}
}
