package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Person struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Name      string        `bson:"name"`
	Phone     string        `bson:"phone"`
	Timestamp time.Time     `bson:"timestamp"`
	FAQList   []*FAQ        `bson:"faqlist"`
}
