package models

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id         bson.ObjectId `bson:"_id,omitempty"`
	FullName   string        `bson:"fullname"`
	Email      string        `bson:"email"`
	TimeZone   string        `bson:"timezone"`
	Language   string        `bson:"language"`
	Password   string        `bson:"password"`
	CourseList []*Training   `bson:"course_list"`
}
