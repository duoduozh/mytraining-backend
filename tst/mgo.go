package main

import (
    "fmt"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type Training struct {
    Id              bson.ObjectId       `bson:"_id,omitempty"` 
    Name            string              `bson:"name"`
    Overview        string              `bson:"overview"`
    BasicInfo       string              `bson:"basic_info"`
    Commitment      string              `bson:"commitment"`
    HowToPass       string              `bson:"how_to_pass"`
    AverageRating   int                 `bson:"average_rating"`
    Icon            string              `bson:"icon"`
    Specification   string              `bson:"specification_info"`
    Forum           string              `bson:"forum"`
    Language        []string            `bson:"language"`
}

type Lecture struct {
    Id              bson.ObjectId       `bson:"_id,omitempty"`
    Name            string              `bson:"name"`
    Organization    string              `bson:"organization"`
}

type Module struct {
    Id              bson.ObjectId       `bson:"_id,omitempty"`
    Week            int                 `bson:"week"`
    Module          string              `bson:"module"`
    Title           string              `bson:"title"`
    Description     string              `bson:"description"`
    Video           int                 `bson:"video"`
    Reading         int                 `bson:"reading"`
    Practice        int                 `bson:"practice"`
    Duration        int                 `bson:"duration"`
    Grade           int                 `bson:"grade"`
}

func main(){
    session, err := mgo.Dial("127.0.0.1:32768")
    if err != nil {
        panic(err)
    } else {
        fmt.Printf("connect success\n")
    }

    defer session.Close()

}
