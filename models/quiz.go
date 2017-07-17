package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Quiz struct {
	Type  string    `bson:"type"`
	Grade string    `bson:"grade"`
	Due   time.Time `bson:"due"`
}
