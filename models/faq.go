package models

import (
//"gopkg.in/mgo.v2/bson"
)

type FAQ struct {
	Question string `bson:"question" json:"question"`
	Answer   string `bson:"answer" json:"answer"`
}
