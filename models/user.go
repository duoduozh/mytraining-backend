package models

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id                     bson.ObjectId `bson:"_id,omitempty"`
	FullName               string        `bson:"fullname"`
	Email                  string        `bson:"email"`
	Language               string        `bson:"language"`
	Password               string        `bson:"password"`
	Birthday               time.Time     `bson:"birthday"`
	CourseList             []*Training   `bson:"course_list,omitempty"`
	RecentViewedCourseList []string      `bson:"recent_viewed_course_list"`
}
