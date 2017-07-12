package models

type Lecture struct {
	Name         string `bson:"name"`
	Description  string `bson:"description"`
	Organization string `bson:"organization"`
}
