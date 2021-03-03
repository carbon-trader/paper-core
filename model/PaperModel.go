package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Paper struct
type Paper struct {
	ID             bson.ObjectId `bson:"_id" json:"id"`
	Paper          string        `bson:"paper" json:"paper"`
	CommercialName string        `bson:"" json:""`
	CompanyName    string        `bson:"" json:""`
	Type           string        `bson:"" json:""`
	Sector         string        `bson:"" json:""`
	SubSector      string        `bson:"" json:""`
	MarketValue    string        `bson:"" json:""`
	CompanyValue   string        `bson:"" json:""`
	NStock         string        `bson:"" json:""`
	LastBalance    time.Time     `bson:"" json:""`
	CreatedAt      time.Time     `bson:"" json:""`
	UpdatedAt      time.Time     `bson:"" json:""`
}
