package main

type StudentsStorage interface {
	Init()
	Add(s Student) error
	Get(key string) (Student, bool)
	GetAll() []Student
	Count() int
} 
/*
{
	"name": <value>,	e.g. "top"
	"age": <value>		e.g. 21
}
 */

type Student struct {
	Name string `json:"name"`
	Age int `json:"age"`
	StudentID string `json:"studentid"`
}

type StudentsDB struct {
	students map[string]Student
}

func (db *StudentsDB)Init()  {
	db.students = make(map[string]Student)
}

func (db *StudentsDB)Add(s Student) error {
	db.students[s.StudentID] = s
	return nil
}

func (db *StudentsDB)Count() int {
	return len(db.students)
}

func (db *StudentsDB)Get(keyId string) (Student, bool){
	s, ok:= db.students[keyId]
	return s, ok
}

func (db *StudentsDB)GetAll() []Student{
	all := make([]Student, 0, db.Count())
	for _, s := range db.students{
		all = append(all, s)
	}
	return all
}