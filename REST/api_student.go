package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"net/http"
)

func replyWithAllStudents(w http.ResponseWriter, db StudentsStorage) {
	if db.Count() == 0 {
		json.NewEncoder(w).Encode([]Student{})
	} else {
		a := make([]Student, 0, db.Count())
		for _, s := range db.GetAll(){
			a = append(a, s)
		}
		json.NewEncoder(w).Encode(a)
	}
}

func replyWithStudent(w http.ResponseWriter, db StudentsStorage, id string)  {
	s, ok := db.Get(id)
	if !ok {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(s)
}

func handlerStudent(w http.ResponseWriter, r *http.Request)  {
	switch r.Method {
	case "POST":
		if r.Body == nil{
			http.Error(w, "Student POST request must have JSON body", http.StatusBadRequest)
			return
		}
		var s Student
		err:=json.NewDecoder(r.Body).Decode(&s)
		if err != nil{
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// check if the student is new
		_, ok :=db.Get(s.StudentID)
		if ok {
			// TODO find a better Error Code (HTTP Status)
			http.Error(w, "Student already exists. Use PUT to modify.", http.StatusBadRequest)
			return
		}
		// new student
		db.Add(s)
		fmt.Fprint(w,"ok") // 200 by default
		return
	case "GET": // student/<id>
		http.Header.Add(w.Header(),"content-type", "application/json")
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) != 3 || parts[1] != "student"{
			http.Error(w, "Malformed URL", http.StatusBadRequest)
			return
		}
		if parts[2] == ""{
			replyWithAllStudents(w, db)
		}else {
			replyWithStudent(w, db, parts[2])
		}
	default:
		http.Error(w, "not implemented yet", http.StatusNotImplemented)
		return
	}
}