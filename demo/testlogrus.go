package demo

import (
    "github.com/onrik/logrus/filename"
    log "github.com/sirupsen/logrus"
    "errors"
)

func TestLog() {
    testLogError()
    //testFileInfo()
    //testInfo()
    //testWarn()
    //testError()
    //testCommonField()
}

func testLogError() {
    err := raiseError()
    log.Errorf("Error happens, below is the error information: \n %v", err)
}

func raiseError() error {
    return errors.New("This is an test error")
}

func testFileInfo() {
    filenameHook := filename.NewHook()
    filenameHook.Field = "filename_line" // Customize source field name
    log.AddHook(filenameHook)
}

func testInfo() {
    log.WithFields(log.Fields{
        "animal": "walrus",
        "size":   10,
    }).Info("A group of walrus emerges from the ocean")
}

func testWarn() {
    log.WithFields(log.Fields{
        "omg":    true,
        "number": 122,
    }).Warn("The group's number increased tremendously!")
}

func testError() {
    log.WithFields(log.Fields{
        "omg":    true,
        "number": 100,
    }).Error("The ice breaks!")
}

func testCommonField() {
    // A common pattern is to re-use fields between logging statements by re-using
    // the logrus.Entry returned from WithFields()
    contextLogger := log.WithFields(log.Fields{
        "common": "this is a common field",
        "other":  "I also should be logged always",
    })

    contextLogger.Info("I'll be logged with common and other field")
    contextLogger.Info("Me too")
}