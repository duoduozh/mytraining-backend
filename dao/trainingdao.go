package dao

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"models"
	"time"
)

type trainingdao struct {
}

func (t *trainingdao) CreateTraining(training model.training) error {
	//	session := global.
}
