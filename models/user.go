package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type User struct {
	Id                     bson.ObjectId `bson:"_id,omitempty"`
	Name                   string        `bson:"name"`
	Email                  string        `bson:"email"`
	Language               string        `bson:"language"`
	Password               string        `bson:"password"`
	Birthday               time.Time     `bson:"birthday"`
	CourseList             []*Training   `bson:"course_list,omitempty"`
	RecentViewedCourseList []string      `bson:"recent_viewed_course_list"`
}
