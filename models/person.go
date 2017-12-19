package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Person struct {
	ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name      string        `bson:"name" json:"name"`
	Phone     string        `bson:"phone" json:"phone"`
	Timestamp time.Time     `bson:"timestamp" json:"timestamp"`
	FAQList   []*FAQ        `bson:"faqlist" json:"faq_list"`
}
