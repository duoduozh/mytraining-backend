package models

type Syllabus struct {
	Week             int    `bson:"week"`
	Module           string `bson:"module"`
	Title            string `bson:"title"`
	Description      string `bson:"description"`
	VideoDuration    int    `bson:"videoDuration"`
	ReadingDuration  int    `bson:"readingDuartion"`
	PracticeDuration int    `bson:"practiceDuartion"`
	Duration         int    `bson:"duration"`
	Grade            string `bson:"grade"`
}
