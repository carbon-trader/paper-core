package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Paper struct
type Paper struct {
	ID             bson.ObjectId `bson:"_Id"            json:"Id"`
	Paper          string        `bson:"Paper"          json:"Paper"`
	CommercialName string        `bson:"CommercialName" json:"CommercialName"`
	CompanyName    string        `bson:"CompanyName"    json:"CompanyName"`
	Type           string        `bson:"Type"           json:"Type"`
	Sector         string        `bson:"Sector"         json:"Sector"`
	SubSector      string        `bson:"SubSector"      json:"SubSector"`
	MarketValue    string        `bson:"MarketValue"    json:"MarketValue"`
	CompanyValue   string        `bson:"CompanyValue"   json:"CompanyValue"`
	NStock         string        `bson:"NStock"         json:"NStock"`
	LastBalance    time.Time     `bson:"LastBalance"    json:"LastBalance"`
	CreatedAt      time.Time     `bson:"CreatedAt"      json:"CreatedAt"`
	UpdatedAt      time.Time     `bson:"UpdatedAt"      json:"UpdatedAt"`
}
