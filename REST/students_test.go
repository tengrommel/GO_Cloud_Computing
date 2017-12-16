package main
import "testing"

func Test_addStudent(t *testing.T) {
	db := StudentsDB{}
	db.Add(Student{"Tom", 21})
	if db.Count() != 1 {
		t.Error("Wrong student count")
	}
	s := db.Get(0)
	if s.Name != "Tom" {
		t.Error("Student Tom was not added.")
	}
}

func Test_multipleStudents(t *testing.T)  {
	testData := []Student{
		Student{"Bob", 21},
		Student{"Alice", 20},
		Student{"ten", 36},
	}
	db := StudentsDB{}
	for _, s := range testData{
		db.Add(s)
	}
	if db.Count()!=len(testData){
		t.Error("Wrong number of students")
	}
	for i, _:= range db.students  {
		if (db.Get(i).Name!=testData[i].Name) {
			t.Error("Wrong name")
		}
		if (db.Get(i).Age!=testData[i].Age) {
			t.Error("Wrong age")
		}
	}


}