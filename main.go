package main

import (
	"encoding/json"
	"fmt"
	"mytraining_backend/config"
	"mytraining_backend/dao"
	"mytraining_backend/models"
	"mytraining_backend/util"
)

func main() {
	// Initialize Environment
	config.LoadConfig()
	err := util.InitDB()
	if err != nil {
		panic(err)
	}
	defer util.CloseDB()

	var training models.Training
	training.Name = `HTML, CSS, and Javascript for Web Developers`
	training.Overview = `Welcome to HTML, CSS, and Javascript for Web Developers! You're joining thousands of learners currently enrolled in the course. I'm excited to have you in the class and look forward to your contributions to the learning community.`
	lecture := &models.Lecture{
		Name:         "Carson, Zhang",
		Description:  "Software Engineer, 10+ years experience",
		Organization: "DXC, previously HPE, and previously previously HP",
	}
	training.LectureList = append(training.LectureList, lecture)
	training.BasicInfo = `Course 4 of 6 in the Ruby on Rails Web Development Specialization.`
	training.Commitment = `5 weeks of study, 4-6 hours/week`
	training.Language = append(training.Language, "English")
	training.HowToPass = `Pass all graded assignments to complete the course.`
	training.AverageRating = 4.5
	training.Icon = `http://via.placeholder.com/350x150`
	training.SpecificationInfo = `TODO`

	syllabus1 := &models.Syllabus{
		Week:             1,
		Module:           `module1`,
		Title:            `Introduction to HTML5`,
		Description:      `In this module we will learn the basics of HTML5. We'll start with instructional videos on how to set up your development environment, go over HTML5 basics like valid document structure, which elements can be included inside other elements and which can not, discuss the meaning and usefulness of HTML5 semantic tags, and go over essential HTML5 tags.`,
		VideoDuration:    60,
		ReadingDuration:  40,
		PracticeDuration: 20,
		Duration:         7,
		Grade:            `quiz`,
	}

	syllabus2 := &models.Syllabus{
		Week:             2,
		Module:           `module2`,
		Title:            `Introduction to HTML5`,
		Description:      `In this module we will learn the basics of HTML5. We'll start with instructional videos on how to set up your development environment, go over HTML5 basics like valid document structure, which elements can be included inside other elements and which can not, discuss the meaning and usefulness of HTML5 semantic tags, and go over essential HTML5 tags.`,
		VideoDuration:    60,
		ReadingDuration:  40,
		PracticeDuration: 20,
		Duration:         7,
		Grade:            `assignment`,
	}

	syllabus3 := &models.Syllabus{
		Week:             3,
		Module:           `module3`,
		Title:            `Introduction to HTML5`,
		Description:      `In this module we will learn the basics of HTML5. We'll start with instructional videos on how to set up your development environment, go over HTML5 basics like valid document structure, which elements can be included inside other elements and which can not, discuss the meaning and usefulness of HTML5 semantic tags, and go over essential HTML5 tags.`,
		VideoDuration:    60,
		ReadingDuration:  40,
		PracticeDuration: 20,
		Duration:         7,
		Grade:            `assignment`,
	}

	syllabus4 := &models.Syllabus{
		Week:             4,
		Module:           `module4`,
		Title:            `Introduction to HTML5`,
		Description:      `In this module we will learn the basics of HTML5. We'll start with instructional videos on how to set up your development environment, go over HTML5 basics like valid document structure, which elements can be included inside other elements and which can not, discuss the meaning and usefulness of HTML5 semantic tags, and go over essential HTML5 tags.`,
		VideoDuration:    60,
		ReadingDuration:  40,
		PracticeDuration: 20,
		Duration:         7,
		Grade:            `assignment`,
	}

	syllabus5 := &models.Syllabus{
		Week:             5,
		Module:           `module5`,
		Title:            `Introduction to HTML5`,
		Description:      `In this module we will learn the basics of HTML5. We'll start with instructional videos on how to set up your development environment, go over HTML5 basics like valid document structure, which elements can be included inside other elements and which can not, discuss the meaning and usefulness of HTML5 semantic tags, and go over essential HTML5 tags.`,
		VideoDuration:    60,
		ReadingDuration:  40,
		PracticeDuration: 20,
		Duration:         7,
		Grade:            `assignment`,
	}

	training.SyllabusList = append(training.SyllabusList, syllabus1, syllabus2, syllabus3, syllabus4, syllabus5)

	resultstring, _ := json.Marshal(training)
	fmt.Printf("training is %v\n\n\n\n\n", string(resultstring))

	err = dao.CreateTraining()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
