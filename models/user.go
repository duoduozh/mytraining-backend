package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type User struct {
	Id                     bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name                   string        `bson:"name" json:"name"`
	Email                  string        `bson:"email" json:"email"`
	Language               string        `bson:"language" json:"language"`
	Password               string        `bson:"password" json:"password"`
	Birthday               time.Time     `bson:"birthday" json:"birthday"`
	CourseList             []*Training   `bson:"course_list,omitempty" json:"course_list,omitempty"`
	RecentViewedCourseList []string      `bson:"recent_viewed_course_list" json:"recent_viewed_course_list"`
}
