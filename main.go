package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"mytraining_backend/config"
	"mytraining_backend/dao"
	//"mytraining_backend/errorhandle"
	"mytraining_backend/models"
	"mytraining_backend/util"
)

func main() {
	// Initialize Database Connection
	config.LoadConfig()
	err := util.InitDB()
	if err != nil {
		panic(err)
	}
	defer util.CloseDB()

	util.InitLog()
	LogTest()
	//dao.DBOperationDemo()
	//InitializeData() // This statement is to initialize some data for database testing, need to remove

	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	//r := gin.New()
	//r.Use(errorhandle.MyRecovery())
	r.GET("/test/user/:user/psw/:psw", TestHandler)
	r.GET("/mytraining/api/v1/user/:name", FindUserByNameHandler)
	r.GET("/mytraining/api/v1/training/name/:name", FindTrainingByNameHandler)
	r.GET("/mytraining/api/v1/training/tag/:tag", FindTrainingByTagHandler)
	r.GET("/mytraining/api/v1/training/lang/:lang", FindTrainingByLanguageHandler)
	r.PUT("/mytraining/api/v1/user", UpdateUserHandler)
	r.Run(":8080")
}

func LogTest() {
	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	//log.WithFields(log.Fields{
	//    "omg":    true,
	//    "number": 100,
	//}).Fatal("The ice breaks!")

	// A common pattern is to re-use fields between logging statements by re-using
	// the logrus.Entry returned from WithFields()
	contextLogger := log.WithFields(log.Fields{
		"common": "this is a common field",
		"other":  "I also should be logged always",
	})

	contextLogger.Info("I'll be logged with common and other field")
	contextLogger.Info("Me too")
}

// This Handler is used to test parameter binding
func TestHandler(c *gin.Context) {
	user := c.Params.ByName("user")
	psw := c.Params.ByName("psw")
	result := fmt.Sprintf("user is %v and psw is %v\n", user, psw)
	c.JSON(200, gin.H{"status": "ok", "Result": result})
}

func FindUserByNameHandler(c *gin.Context) {
	userName := c.Params.ByName("name")
	user, err := dao.FindUserByName(userName)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	if &user == nil {
		fmt.Println("No user is found")
		c.JSON(200, gin.H{"user": user, "status": "no value"})
	} else {
		jsonUserString, _ := json.Marshal(user)
		fmt.Printf("FindUserByName result is: %v\n", string(jsonUserString))
		c.JSON(200, gin.H{"user": user, "status": "ok"})
	}
}

func UpdateUserHandler(c *gin.Context) {
	var user models.User
	err := c.Bind(&user)
	if err == nil {
		jsonUserString, _ := json.Marshal(user)
		fmt.Printf("Received User Json is: %v\n", string(jsonUserString))
		dao.UpdateUser(user)
		c.JSON(200, gin.H{"user": string(jsonUserString), "status": "ok"})
	} else {
		c.JSON(400, gin.H{"status": "400 Bad Request", "error": err.Error()})
	}
}

func FindTrainingByLanguageHandler(c *gin.Context) {
	trainingLanguage := c.Params.ByName("lang")
	training, err := dao.FindTrainingByLanguage(trainingLanguage)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	if training == nil {
		fmt.Println("No training is found")
		c.JSON(200, gin.H{"training": training, "status": "no value"})
	} else {
		jsonTrainingString, _ := json.Marshal(training)
		fmt.Printf("FindTrainingByLanguage success, result is: %v\n", string(jsonTrainingString))
		c.JSON(200, gin.H{"training": training, "status": "ok"})
	}
}

func FindTrainingByTagHandler(c *gin.Context) {
	trainingTag := c.Params.ByName("tag")
	training, err := dao.FindTrainingByTag(trainingTag)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	if training == nil {
		fmt.Println("No training is found")
		c.JSON(200, gin.H{"training": training, "status": "no value"})
	} else {
		jsonTrainingString, _ := json.Marshal(training)
		fmt.Printf("FindTrainingByTag success, result is: %v\n", string(jsonTrainingString))
		c.JSON(200, gin.H{"training": training, "status": "ok"})
	}
}

func FindTrainingByNameHandler(c *gin.Context) {
	trainingName := c.Params.ByName("name")
	training, err := dao.FindTrainingByName(trainingName)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	if training == nil {
		fmt.Println("No training is found")
		c.JSON(200, gin.H{"training": training, "status": "no value"})
	} else {
		jsonTrainingString, _ := json.Marshal(training)
		fmt.Printf("FindTrainingByName success, result is: %v\n", string(jsonTrainingString))
		c.JSON(200, gin.H{"training": training, "status": "ok"})
	}
}