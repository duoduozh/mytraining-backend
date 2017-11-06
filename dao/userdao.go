package dao

import (
	"errors"
	"fmt"
	"github.com/alioygur/is"
	"gopkg.in/mgo.v2/bson"
	"mytraining_backend/models"
	"mytraining_backend/util"
)

func CreateUser(user models.User) error {
	session, err := util.GetDBSession()
	if err != nil {
		fmt.Println(err)
		panic(nil)
	}
	defer session.Close()

	c := session.DB("").C("user")
	err = c.Insert(user)
	if err != nil {
		panic(err)
	}

	return nil
}

func FindUserByEmail(email string) (user *models.User, err error) {
	if len(email) <= 0 {
		return user, errors.New("Email is null or empty")
	}

	if !(is.Email(email)) {
		return user, errors.New("Param is not a valid email: " + email)
	}

	session, err := util.GetDBSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("").C("user")
	var userList []models.User
	err = c.Find(bson.M{"email": email}).All(&userList)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	if len(userList) > 0 {
		user = &userList[0]
	}

	return user, nil
}
