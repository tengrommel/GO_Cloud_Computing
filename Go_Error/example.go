package main

import (
	"errors"
	"fmt"
)

var (
	ErrTooOld = errors.New("student too old")
	ErrDuplicateEntry = errors.New("duplicate entry")
)

type Student struct {
	Name string
	Age int
}

var db map[string]Student

func addStudent(s Student) error {
	if s.Age > 30{
		return ErrTooOld
		// error that indicates that student is too old
	}

	// if student with a given name is already stored
	// error indicated duplicate entry
	if _, ok := db[s.Name]; ok{
		return ErrDuplicateEntry
	}
	// store the student
	db[s.Name] = s
	return nil
}

func main()  {
	db = make(map[string]Student)

	tom := Student{"Tom", 40}
	err := addStudent(tom)
	fmt.Println("We got an error: " + err.Error())

	alice := Student{"Alice", 20}
	err = addStudent(alice)

	switch err {
	case ErrTooOld:
		fmt.Println("Tom's too old")
	case ErrDuplicateEntry:
		fmt.Println("Tom is already in DB")
	}

	err = addStudent(alice)
	switch err {
	case ErrTooOld:
		fmt.Println("Tom's too old")
	case ErrDuplicateEntry:
		fmt.Println("Tom is already in DB")
	}
	// handle error depending on what it is
}