package dao

import (
	"errors"
	"github.com/alioygur/is"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
	"mytraining_backend/models"
	"mytraining_backend/util"
)

func CreateUser(user models.User) error {
	session, err := util.GetDBSession()
	if err != nil {
		log.Error(err)
		panic(nil)
	}
	defer session.Close()

	c := session.DB("").C("user")
	err = c.Insert(user)
	if err != nil {
		log.Error(err)
		panic(err)
	}

	return nil
}

func UpdateUser(user models.User) error {
	if &user == nil {
		errors.New("parameter null error, param name: user")
	}

	session, err := util.GetDBSession()
	if err != nil {
		log.Error(err)
		panic(nil)
	}
	defer session.Close()
	c := session.DB("").C("user")
	err = c.Update(bson.M{"_id": user.Id}, user)
	if err != nil {
		log.Error(err)
		panic(err)
	} else {
		return nil
	}
}

func FindUserByEmail(email string) (user models.User, err error) {
	if len(email) <= 0 {
		return user, errors.New("Email is null or empty")
	}

	if !(is.Email(email)) {
		return user, errors.New("Param is not a valid email: " + email)
	}

	session, err := util.GetDBSession()
	if err != nil {
		log.Error(err)
		panic(err)
	}
	defer session.Close()

	c := session.DB("").C("user")
	var userList []models.User
	err = c.Find(bson.M{"email": email}).All(&userList)
	if err != nil {
		log.Error(err)
		panic(err)
	}

	if len(userList) > 0 {
		user = userList[0]
	}

	return user, nil
}

func FindUserByName(name string) (user models.User, err error) {
	if len(name) <= 0 {
		return user, errors.New("Name is null or empty")
	}

	session, err := util.GetDBSession()
	if err != nil {
        log.Error(err)
		panic(err)
	}
	defer session.Close()

	c := session.DB("").C("user")
	var userList []models.User
	err = c.Find(bson.M{"name": name}).All(&userList)
	if err != nil {
        log.Error(err)
		panic(err)
	}

	if len(userList) > 0 {
		user = userList[0]
	}

	return user, nil
}