package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Training struct {
	Id                bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name              string        `bson:"name" json:"name"`
	Overview          string        `bson:"overview" json:"overview"`
	LectureList       []*Lecture    `bson:"lecture_list" json:"lecture_list"`
	BasicInfo         string        `bson:"basic_info" json:"basic_info"`
	Commitment        string        `bson:"commitment" json:"commitment"`
	Language          []string      `bson:"language" json:"language"`
	HowToPass         string        `bson:"how_to_pass" json:"how_to_pass"`
	AverageRating     float32       `bson:"average_rating" json:"average_rating"`
	Icon              string        `bson:"icon" json:"icon"`
	SpecificationInfo string        `bson:"specification_info" json:"specification_info"`
	SyllabusList      []*Syllabus   `bson:"syllabus_list" json:"syllabus_list"`
	FAQList           []*FAQ        `bson:"FAQ_list" json:"faq_list"`
	Forum             string        `bson:"forum" json:"forum"`
	ResourceList      []string      `bson:"resource_list" json:"resource_list"`
	TagList           []string      `bson:"tag_list" json:"tag_list"`
}
