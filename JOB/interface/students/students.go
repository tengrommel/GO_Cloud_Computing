package main

import "fmt"

type Student struct {
	name string
	id int
}

type ClassRoom struct {
	students []Student
}

func (c *ClassRoom)add(name string, id int)  {
	c.students = append(c.students, Student{name, id})
}

func (c *ClassRoom)list() []Student {
	return c.students
}

var classrooms map[string]*ClassRoom

func main() {
	c := ClassRoom{}
	fmt.Println(c)

}
