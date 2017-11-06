package dao

import (
	"errors"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"mytraining_backend/models"
	"mytraining_backend/util"
)

func CreateTraining(training models.Training) error {
	session, err := util.GetDBSession()
	if err != nil {
		fmt.Println(err)
		panic(nil)
	}
	defer session.Close()

	c := session.DB("").C("training")
	err = c.Insert(training)
	if err != nil {
		panic(err)
	}

	return nil
}

func FindTrainingByName(name string) (trainingList []models.Training, err error) {
	if len(name) <= 0 {
		return trainingList, errors.New("Training Name is null or empty")
	}

	session, err := util.GetDBSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("").C("training")
	err = c.Find(bson.M{"name": name}).All(&trainingList)
	if err != nil {
		panic(err)
	}

	return trainingList, nil
}

func FindTrainingByTag(tag string) (trainingList []models.Training, err error) {
	if len(tag) <= 0 {
		return trainingList, errors.New("Training Name is null or empty")
	}

	session, err := util.GetDBSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("").C("training")
	err = c.Find(bson.M{"tag_list": tag}).All(&trainingList)
	if err != nil {
		panic(err)
	}

	return trainingList, nil
}

func FindTrainingByLanguage(language string) (trainingList []models.Training, err error) {
	if len(language) <= 0 {
		language = "english"
	}

	session, err := util.GetDBSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("").C("training")
	err = c.Find(bson.M{"language": bson.M{"$regex": bson.RegEx{language, "i"}}}).All(&trainingList)
	if err != nil {
		panic(err)
	}

	return trainingList, nil
}