package main

import (
    "fmt"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "time"
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

type Person struct {
    ID              bson.ObjectId       `bson:"_id,omitempty"`
    Name            string
    Phone           string
    Timestamp       time.Time
}

var ( 
    IsDrop = true
)
func main(){
    session, err := mgo.Dial("127.0.0.1")
    
    if err != nil {
        panic(err)
    } else {
        fmt.Printf("connect success\n")
    }

    defer session.Close()

    session.SetMode(mgo.Monotonic, true)

    // Drop DB
    if IsDrop {
        err = session.DB("test").DropDatabase()
        if err != nil {
            panic(err)
        }
    } 

    // Collection People
    c := session.DB("test").C("people")

    // Index
    index := mgo.Index{
        Key: []string{"name", "phone"},
        Unique: true,
        DropDups: true,
        Background: true,
        Sparse: true,
    }

    err = c.EnsureIndex(index)
    if err != nil {
        panic(err)
    }

    // Insert Datas
    err = c.Insert(&Person{Name: "Ale", Phone: "+55 53 1234 4321", Timestamp: time.Now()},
                    &Person{Name: "Cla", Phone: "+66 33 1234 5678", Timestamp: time.Now()})

    if err != nil {
        panic(err)
    }

    // Query One
    result := Person{}
    err = c.Find(bson.M{"name": "Ale"}).Select(bson.M{"phone": 0}).One(&result)
    if err != nil {
        panic(err)
    }
    fmt.Println("Result Person", result)

    // Query All
    var resultPersonList []Person
    err = c.Find(bson.M{"name": "Ale"}).Sort("-timestamp").All(&resultPersonList)
    if err != nil {
        panic(err)
    }
    fmt.Println("Result Person List", resultPersonList)

}
