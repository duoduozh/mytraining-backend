package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Module struct {
	Id          bson.ObjectId `bson:"_id,omitempty"`
	Week        int           `bson:"week"`
	Module      string        `bson:"module"`
	Title       string        `bson:"title"`
	Description string        `bson:"description"`
	Video       int           `bson:"video"`
	Reading     int           `bson:"reading"`
	Practice    int           `bson:"practice"`
	Duration    int           `bson:"duration"`
	Grade       int           `bson:"grade"`
}
