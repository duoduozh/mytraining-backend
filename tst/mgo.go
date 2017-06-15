package main

import (
    "fmt",
    "log"
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
}

func main(){

}

