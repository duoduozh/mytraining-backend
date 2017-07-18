package models

type Syllabus struct {
	Week             int     `bson:"week"`
	Module           string  `bson:"module"`
	Title            string  `bson:"title"`
	Description      string  `bson:"description"`
	VideoDuration    int     `bson:"videoDuration"`
	VideoProgress    float32 `bson:"videoProgress,omitempty"`
	ReadingDuration  int     `bson:"readingDuartion"`
	ReadingProgress  float32 `bson:"readingProgress,omitempty"`
	PracticeDuration int     `bson:"practiceDuartion"`
	PracticeProgress float32 `bson:"practiceProgress,omitempty"`
	Duration         int     `bson:"duration"`
	Graded           string  `bson:"graded"`
}
