package dao

import (
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"mytraining_backend/models"
	"mytraining_backend/util"
	"time"
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

func DBOperationDemo() error {
	session, err := util.GetDBSession()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer session.Close()

	// Collection People
	c := session.DB("").C("people")

	person1 := models.Person{
		Name:      "Ale",
		Phone:     "+55 53 1234 4321",
		Timestamp: time.Now(),
		FAQList: []*models.FAQ{
			&models.FAQ{Question: "q1", Answer: "a1"},
			&models.FAQ{Question: "q2", Answer: "a2"},
		},
	}

	person2 := models.Person{
		Name:      "Cla",
		Phone:     "+66 33 1234 5678",
		Timestamp: time.Now(),
	}

	// Insert Datas
	err = c.Insert(&person1, &person2)

	if err != nil {
		panic(err)
	}

	// Query One
	result := models.Person{}
	err = c.Find(bson.M{"name": "Ale"}).Select(bson.M{"phone": 0}).One(&result)
	if err != nil {
		panic(err)
	}
	resultstring, _ := json.Marshal(&result)
	fmt.Printf("Result Person %v\n", string(resultstring))

	// Query All
	var resultPersonList []models.Person
	err = c.Find(bson.M{"name": "Ale"}).Sort("-timestamp").All(&resultPersonList)
	if err != nil {
		panic(err)
	}
	resultstring, _ = json.Marshal(&resultPersonList)
	fmt.Printf("Result Person List %v\n", string(resultstring))

	// Update
	colQuerier := bson.M{"name": "Ale"}
	change := bson.M{"$set": bson.M{"phone": "+86 99 8888 7773", "timestamp": time.Now()}}
	err = c.Update(colQuerier, change)
	if err != nil {
		panic(err)
	}

	// Query All
	err = c.Find(bson.M{"name": "Ale"}).Sort("-timestamp").All(&resultPersonList)
	if err != nil {
		panic(err)
	}
	resultstring, _ = json.Marshal(&resultPersonList)
	fmt.Printf("Results All: %v\n", string(resultstring))

	return nil
}
