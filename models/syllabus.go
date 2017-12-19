package models

type Syllabus struct {
	Week             int     `bson:"week" json:"week"`
	Module           string  `bson:"module" json:"module"`
	Title            string  `bson:"title" json:"title"`
	Description      string  `bson:"description" json:"description"`
	VideoDuration    int     `bson:"videoDuration" json:"video_duration"`
	VideoProgress    float32 `bson:"videoProgress,omitempty" json:"video_progress"`
	ReadingDuration  int     `bson:"readingDuartion" json:"reading_duration"`
	ReadingProgress  float32 `bson:"readingProgress,omitempty" json:"reading_progress"`
	PracticeDuration int     `bson:"practiceDuartion" json:"practice_duration"`
	PracticeProgress float32 `bson:"practiceProgress,omitempty" json:"practice_progress"`
	Duration         int     `bson:"duration" json:"duration"`
	Graded           string  `bson:"graded" json:"graded"`
}
