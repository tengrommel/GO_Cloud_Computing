package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

// StudentsMongoDB stores the details of the DB connection.
type StudentsMongoDB struct {
	DatabaseURL string
	DatabaseName string
	StudentsCollectionName string
}

/*
Init initializes the mongo storage.
 */
func (db *StudentsMongoDB)Init()  {
	session, err := mgo.Dial(db.DatabaseURL)
	if err!=nil{
		panic(err)
	}
	defer session.Close()

	index :=mgo.Index{
		Key: []string{"studentid"},
		Unique:true,
		DropDups:true,
		Background:true,
		Sparse:true,
	}
	err = session.DB(db.DatabaseName).C(db.StudentsCollectionName).EnsureIndex(index)
	if err != nil{
		panic(err)
	}
}

/*
Add adds new students to the storage
 */
func (db *StudentsMongoDB)Add(s Student) error {
	session, err := mgo.Dial(db.DatabaseURL)
	if err!=nil{
		panic(err)
	}
	defer session.Close()

	err = session.DB(db.DatabaseName).C(db.StudentsCollectionName).Insert(s)
	if err!= nil{
		fmt.Printf("error in Insert(): %v", err.Error())
		return err
	}
	return nil
}

/*
Count returns the current count of the students in in-memory storage.
 */
func (db *StudentsMongoDB)Count() int {
	session, err := mgo.Dial(db.DatabaseURL)
	if err!=nil{
		panic(err)
	}
	defer session.Close()

	count, err := session.DB(db.DatabaseName).C(db.StudentsCollectionName).Count()
	if err!=nil{
		fmt.Printf("error in Count(): %v", err.Error())
		return -1
	}
	return count
}

/*
Get returns a students with a given ID or empty student struct.
 */
func (db *StudentsMongoDB)Get(keyId string) (Student, bool){
	session, err := mgo.Dial(db.DatabaseURL)
	if err!=nil{
		panic(err)
	}
	defer session.Close()

	student := Student{}
	allWasGood := true
	err = session.DB(db.DatabaseName).C(db.StudentsCollectionName).Find(bson.M{"studentid": keyId}).One(&student)
	if err != nil{
		allWasGood = false
	}
	return student, allWasGood
}

/*
Get returns a students with a given ID or empty student struct.
 */
func (db *StudentsMongoDB)GetAll() []Student{
	session, err := mgo.Dial(db.DatabaseURL)
	if err!=nil{
		panic(err)
	}
	defer session.Close()

	var all []Student

	err = session.DB(db.DatabaseName).C(db.StudentsCollectionName).Find(bson.M{}).All(&all)
	if err != nil{
		return []Student{}
	}
	return all
}