package main

/*
{
	"name": <value>,	e.g. "top"
	"age": <value>		e.g. 21
}
 */

type Student struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Id string `json:"id"`
}

type StudentsDB struct {
	students map[string]Student
}

func (db *StudentsDB)Init()  {
	db.students = make(map[string]Student)
}

func (db *StudentsDB)Add(s Student)  {
	db.students[s.Id] = s
}

func (db *StudentsDB)Count() int {
	return len(db.students)
}

func (db *StudentsDB)Get(keyId string) (Student, bool){
	s, ok:= db.students[keyId]
	return s, ok
}
