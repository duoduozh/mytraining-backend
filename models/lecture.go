package models

type Lecture struct {
	Name         string `bson:"name" json:"name"`
	Description  string `bson:"description" json:"description"`
	Organization string `bson:"organization" json:"organization"`
}