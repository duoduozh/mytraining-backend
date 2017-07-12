package models

type Syllabus struct {
	Week             string `bson:"week"`
	Module           string `bson:"module"`
	Title            string `bson:"title"`
	Description      string `bson:"description"`
	VideoDuration    int    `bson:"videoDuration"`
	ReadingDuration  int    `bson:"readingDuartion"`
	PracticeDuration int    `bson:"practiceDuartion"`
	Duration         string `bson:"duration"`
	Grade            string `bson:"grade"`
}
