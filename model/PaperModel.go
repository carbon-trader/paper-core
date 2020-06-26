package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//PaperModel struct
type PaperModel struct {
	ID                   bson.ObjectId `bson:"_id" json:"id"`
	Paper                string        `bson:"paper" json:"paper"`
	Type                 string        `bson:"type" json:"type"`
	Company              string        `bson:"company" json:"company"`
	Sector               string        `bson:"sector" json:"sector"`
	Subsector            string        `bson:"subsector" json:"subsector"`
	LastBalanceProcessed time.Time     `bson:"" json:""`
	CreatedAt            time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt            time.Time     `bson:"updated_at" json:"updated_at"`
	CreatedBy            string        `bson:"created_by" json:"created_by"`
	UpdatedBY            string        `bson:"updated_by" json:"updated_by"`
}
