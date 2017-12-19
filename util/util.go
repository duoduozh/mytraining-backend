package util

import (
	"errors"
	"time"
	"gopkg.in/mgo.v2"
	log "github.com/sirupsen/logrus"
    "github.com/onrik/logrus/filename"
	"mytraining_backend/config"
	"os"
)

var dropDatabase bool = false
var dbsession *mgo.Session

func InitializeEnvironment() error {
    err := InitLog()
    if err != nil {
        panic(err)
    }

    config.LoadConfig()
    err = InitDB()
    if err != nil {
        panic(err)
    }

    return nil
}

func InitLog() error {
    filenameHook := filename.NewHook()
    filenameHook.Field = "filename_line" // Customize source field name
    log.AddHook(filenameHook)

	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{ FullTimestamp:true })
    log.Info("Time Stamp Enabled")

	// Output to stdout instead of the default stderr
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	return nil
}

func InitDB() error {
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
			log.Error("database connection failed")
			panic(err)
		}
	}

	if err != nil {
		log.Errorf("Cannot dial to database: %v", config.DatabaseURL)
		panic(err)
	} else {
	    log.Infof("dial to database: %v SUCCESS\n", config.DatabaseURL)
	}

	dbsession.SetMode(mgo.Monotonic, true)

	return nil
}

func CloseDB() error {
	if dbsession != nil {
		dbsession.Close()
		return nil
	} else {
		log.Errorf("db session is nil \n")
		return nil
	}
}

func GetDBSession() (*mgo.Session, error) {
	if dbsession == nil {
		return nil, errors.New("The global mgo session is nil, please initialize the global mgo session before access it")
	}
	return dbsession.Copy(), nil
}