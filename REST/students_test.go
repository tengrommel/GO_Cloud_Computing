package main
import "testing"

func Test_addStudent(t *testing.T) {
	db := StudentsDB{}
	sData := Student{"Tom", 21, "id1"}
	db.Init()
	db.Add(sData)
	if db.Count() != 1 {
		t.Error("Wrong student count")
	}
	s, _ := db.Get(sData.StudentID)
	if s.Name != "Tom" {
		t.Error("Student Tom was not added.")
	}
}

func Test_multipleStudents(t *testing.T)  {
	testData := map[string]Student{
		"id0":{"Bob", 21, "id0"},
		"id1":{"Alice", 20, "id1"},
		"id2":{"ten", 36, "id2"},
	}
	db := StudentsDB{}
	db.Init()
	for _, s := range testData{
		db.Add(s)
	}
	if db.Count()!=len(testData){
		t.Error("Wrong number of students")
	}
	for key := range db.students  {
		s , _:= db.Get(key)
		sTest, _ := testData[key]
		if s.Name!=sTest.Name {
			t.Error("Wrong name")
		}
		if s.Age!=sTest.Age{
			t.Error("Wrong Age")
		}
		if s.StudentID!=sTest.StudentID {
			t.Error("Wrong Id")
		}
	}


}