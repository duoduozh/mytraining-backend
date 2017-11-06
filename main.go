package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"mytraining_backend/config"
	"mytraining_backend/dao"
	"mytraining_backend/models"
	"mytraining_backend/util"
	"time"
)

func InitializeTraining() {
	var err error
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
		Graded:           `quiz`,
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
		Graded:           `assignment`,
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
		Graded:           `assignment`,
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
		Graded:           `assignment`,
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
		Graded:           `assignment`,
	}

	training.SyllabusList = append(training.SyllabusList, syllabus1, syllabus2, syllabus3, syllabus4, syllabus5)

	training.FAQList = []*models.FAQ{
		&models.FAQ{
			Question: `When will I have access to the lectures and assignments?`,
			Answer:   `Once you enroll for a Certificate, you'll have access to all videos, quizzes, and programming assignments (if applicable). Peer review assignments can only be submitted and reviewed once your session has begun. If you choose to explore the cours    e without subscribing, you may not be able to access certain assignments.`,
		},
		&models.FAQ{
			Question: `What will I get if I pay for this course?`,
			Answer:   `If you pay for this course, you will have access to all of the features and content you need to earn a Course Certificate. If you complete the course successfully, your electronic Certificate will be added to your Accomplishments page - from the    re, you can print your Certificate or add it to your LinkedIn profile. Note that the Course Certificate does not represent official academic credit from the partner institution offering the course.`,
		},
	}

	training.Forum = `https://www.coursera.org/learn/html-css-javascript-for-web-developers/discussions`
	training.ResourceList = []string{
		`https://developer.mozilla.org/en-US/docs/Web/JavaScript`,
		`https://www.w3schools.com/css/`,
	}
	training.TagList = []string{"html", "css", "javascript", "web"}

	//resultstring, _ := json.Marshal(training)
	//fmt.Printf("training is %v\n\n\n\n\n", string(resultstring))

	err = dao.CreateTraining(training)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func InitializeUser() {
	var err error
	var user models.User
	user.Name = "carson"
	user.Email = "myzn007@gmail.com"
	user.Language = "chinese"
	user.Password = "123456"
	user.Birthday = time.Date(1983, time.April, 9, 0, 0, 0, 0, time.UTC)
	user.RecentViewedCourseList = []string{}

	//resultstring, _ := json.Marshal(training)
	//fmt.Printf("training is %v\n\n\n\n\n", string(resultstring))

	err = dao.CreateUser(user)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func InitializeData() {
	InitializeTraining()
	InitializeUser()
}

func main() {
	// Initialize Environment
	config.LoadConfig()
	err := util.InitDB()
	if err != nil {
		panic(err)
	}
	defer util.CloseDB()
	// Insert the training data, will be move to unit test file later.
	InitializeData()
	//var trainingList []models.Training
	//trainingList, err = dao.FindTrainingByLanguage("english")
	//if err != nil {
	//	fmt.Println(err)
	//	panic(err)
	//}
	//jsonTrainingListString, _ := json.Marshal(trainingList)

	r := gin.Default()
	r.GET("/user/:name", func(c *gin.Context) {
		userName := c.Params.ByName("name")
		user, err := dao.FindUserByName(userName)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		if user == nil {
			fmt.Println("No user is found")
			c.JSON(200, gin.H{"user": user, "status": "no value"})
		} else {
			jsonUserString, _ := json.Marshal(*user)
			fmt.Printf("FindUserByName result is: %v\n", string(jsonUserString))
			c.JSON(200, gin.H{"user": user, "status": "ok"})
		}
	})
	r.Run(":8080")
	//var user *models.User
	//user, err = dao.FindUserByEmail("myzn007@gmail.com")
}
