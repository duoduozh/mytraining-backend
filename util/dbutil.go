package util

import (
	"errors"
	"fmt"
	mgo "gopkg.in/mgo.v2"
	config "mytraining_backend/config"
	"time"
)

var dropDatabase bool = true
var dbsession *mgo.Session

func Init() error {
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{config.DatabaseURL},
		Timeout:  60 * time.Second,
		Database: config.DatabaseName,
		Username: config.DatabaseUsername,
		Password: config.DatabasePassword,
	}

	var err error
	dbsession, err = mgo.DialWithInfo(mongoDBDialInfo)

	if dropDatabase {
		err = dbsession.DB("").DropDatabase()
		if err != nil {
			//TODO: log.Fatal("Dropdatabase failed")
			panic(err)
		}
	}

	if err != nil {
		//TODO: log.Fatal("Cannot Dial Mongo: ", err)
		panic(err)
	} else {
		fmt.Printf("connect to database: %v SUCCESS", config.DatabaseURL)
	}

	dbsession.SetMode(mgo.Monotonic, true)
	return nil
}

func Close() error {
	if dbsession != nil {
		dbsession.Close()
		return nil
	} else {
		//return errors.New("dbsession is nil when closing the session")
		fmt.Printf("db session is nil \n")
		return nil
	}
}

func GetSession() (*mgo.Session, error) {
	if dbsession == nil {
		return dbsession, errors.New("The global mgo session is nil, please initialize the global mgo session before access it")
	}

	return dbsession.Copy(), nil
}
