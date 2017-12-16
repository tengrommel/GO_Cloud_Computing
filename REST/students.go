package main

type StudentsDB struct {
	students []Student
}

func (db *StudentsDB)Add(s Student)  {
	db.students=append(db.students, s)
}

func (db *StudentsDB)Count() int {
	return len(db.students)
}

func (db *StudentsDB)Get(i int) Student {
	if i < 0 || i >= len(db.students){
		// log and notify
		// throw runtime exception
		return Student{}
	}
	return db.students[i]
}
