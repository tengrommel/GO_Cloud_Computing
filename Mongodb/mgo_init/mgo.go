package main

import (
	"gopkg.in/mgo.v2"
	"log"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

type Person struct {
	Name string
	Phone string
}

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err!=nil{
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("people")
	err = c.Insert(&Person{"teng", "1234567"}, &Person{"David", "1565454"})
	if err !=nil{
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "teng"}).One(&result)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Name:", result.Name)
	fmt.Println("Phone:", result.Phone)
}
