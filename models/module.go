package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Module struct {
	Id          bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Week        int           `bson:"week" json:"week"`
	Module      string        `bson:"module" json:"module"`
	Title       string        `bson:"title" json:"title"`
	Description string        `bson:"description" json:"description"`
	Video       int           `bson:"video" json:"video"`
	Reading     int           `bson:"reading" json:"reading"`
	Practice    int           `bson:"practice" json:"practice"`
	Duration    int           `bson:"duration" json:"duration"`
	Grade       int           `bson:"grade" json:"grade"`
}
