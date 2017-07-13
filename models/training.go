package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Training struct {
	Id                bson.ObjectId `bson:"_id,omitempty"`
	Name              string        `bson:"name"`
	Overview          string        `bson:"overview"`
	LectureList       []*Lecture    `bson:"lecture_list"`
	BasicInfo         string        `bson:"basic_info"`
	Commitment        string        `bson:"commitment"`
	Language          []string      `bson:"language"`
	HowToPass         string        `bson:"how_to_pass"`
	AverageRating     float32       `bson:"average_rating"`
	Icon              string        `bson:"icon"`
	SpecificationInfo string        `bson:"specification_info"`
	SyllabusList      []*Syllabus   `bson:"syllabus_list"`
	FAQList           []*FAQ        `bson:"FAQ_list"`
	Forum             string        `bson:"forum"`
	ResourceList      []string      `bson:"resource_list"`
	TagList           []string      `bson:"tag_list"`
}
