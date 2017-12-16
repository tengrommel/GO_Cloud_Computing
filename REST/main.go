package main

import (
	"net/http"
	"fmt"
	"strings"
	"encoding/json"
)


func handlerHello(w http.ResponseWriter, r *http.Request)  {
	parts := strings.Split(r.URL.Path, "/")
	if (len(parts) != 4){
		status := 400
		http.Error(w, http.StatusText(status), status)
		return
	}
	name := parts[2]
	fmt.Fprintln(w, parts)
	fmt.Fprintf(w, "Hello %s %s!\n", name, parts[3])
}

func replyWithAllStudents(w http.ResponseWriter, db *StudentsDB) {
	if db.students == nil {
		json.NewEncoder(w).Encode([]Student{})
	} else {
		json.NewEncoder(w).Encode(db.students)
	}
}

func replyWithStudent(w http.ResponseWriter, db *StudentsDB, id string)  {
	s, ok := db.Get(id)
	if !ok {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(s)
}

func handlerStudent(w http.ResponseWriter, r *http.Request)  {
	//
	db := StudentsDB{}
	db.Init()
	db.Add(Student{"tom", 21, "id0"})
	if r.Method == "POST"{
		http.Error(w, "not implemented yet", http.StatusNotImplemented)
		return
	}
	//
	http.Header.Add(w.Header(),"content-type", "application/json")
	parts := strings.Split(r.URL.Path, "/")

	if len(parts) != 3{
		return
	}

	if parts[2] == ""{
		replyWithAllStudents(w, &db)
		}else {
			replyWithStudent(w, &db, parts[2])
		}
}

func main() {
	http.HandleFunc("/hello/", handlerHello)
	http.HandleFunc("/student/", handlerStudent)
	http.ListenAndServe(":8000", nil)
}
