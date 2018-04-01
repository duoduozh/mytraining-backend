package main

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"mytraining_backend/config"
	"mytraining_backend/dao"
	"mytraining_backend/errorhandle"
	"mytraining_backend/models"
	"mytraining_backend/util"
	"os"
	"os/signal"
	"time"
)

func main() {
	err := util.InitializeEnvironment()
	if err != nil {
		panic(err)
	}
	defer util.CloseDB() // Close DB connection before quit main method

	go func() {
		gin.SetMode(gin.DebugMode)
		r := gin.New()
		r.Use(errorhandle.MyRecovery())
		r.GET(config.API_Prefix+"/training/name/:name", FindTrainingByNameHandler)
		r.GET(config.API_Prefix+"/training/tag/:tag", FindTrainingByTagHandler)
		r.GET(config.API_Prefix+"/training/lang/:lang", FindTrainingByLanguageHandler)
		r.Run(":8080")
	}()

	go func() {
		gin.SetMode(gin.DebugMode)
		r := gin.New()
		r.Use(errorhandle.MyRecovery())
		r.GET(config.API_Prefix+"/user/:name", FindUserByNameHandler)
		r.PUT(config.API_Prefix+"/user", UpdateUserHandler)
		r.Run(":8081")
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	log.Println("Server exiting")
}

func FindUserByNameHandler(c *gin.Context) {
	userName := c.Params.ByName("name")
	user, err := dao.FindUserByName(userName)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	if &user == nil {
		log.Info("No user is found with name %v", userName)
		c.JSON(200, gin.H{"user": user, "status": "no value"})
	} else {
		jsonUserString, _ := json.Marshal(user)
		log.Infof("FindUserByName result is: %v\n", string(jsonUserString))
		c.JSON(200, gin.H{"user": user, "status": "ok"})
	}
}

func UpdateUserHandler(c *gin.Context) {
	var user models.User
	err := c.Bind(&user)
	if err == nil {
		jsonUserString, _ := json.Marshal(user)
		log.Infof("Received User Json is: %v\n", string(jsonUserString))
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
		log.Error(err)
		panic(err)
	}
	if training == nil {
		log.Info("No training is found")
		c.JSON(200, gin.H{"training": training, "status": "no value"})
	} else {
		jsonTrainingString, _ := json.Marshal(training)
		log.Infof("FindTrainingByLanguage success, result is: %v\n", string(jsonTrainingString))
		c.JSON(200, gin.H{"training": training, "status": "ok"})
	}
}

func FindTrainingByTagHandler(c *gin.Context) {
	trainingTag := c.Params.ByName("tag")
	training, err := dao.FindTrainingByTag(trainingTag)
	if err != nil {
		log.Info(err)
		panic(err)
	}
	if training == nil {
		log.Info("No training is found")
		c.JSON(200, gin.H{"training": training, "status": "no value"})
	} else {
		jsonTrainingString, _ := json.Marshal(training)
		log.Infof("FindTrainingByTag success, result is: %v\n", string(jsonTrainingString))
		c.JSON(200, gin.H{"training": training, "status": "ok"})
	}
}

func FindTrainingByNameHandler(c *gin.Context) {
	trainingName := c.Params.ByName("name")
	training, err := dao.FindTrainingByName(trainingName)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	if training == nil {
		log.Info("No training is found")
		c.JSON(200, gin.H{"training": training, "status": "no value"})
	} else {
		jsonTrainingString, _ := json.Marshal(training)
		log.Infof("FindTrainingByName success, result is: %v\n", string(jsonTrainingString))
		c.JSON(200, gin.H{"training": training, "status": "ok"})
	}
}
