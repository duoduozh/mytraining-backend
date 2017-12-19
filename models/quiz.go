package models

import (
	//	"gopkg.in/mgo.v2/bson"
	"time"
)

type Quiz struct {
	Type  string    `bson:"type" json:"type"`
	Grade string    `bson:"grade" json:"grade"`
	Due   time.Time `bson:"due" json:"due"`
}
