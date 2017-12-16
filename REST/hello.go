package main

import (
	"net/http"
	"fmt"
	"strings"
	//"encoding/json"
	"encoding/json"
	"strconv"
)

/*
{
	"name": <value>,	e.g. "top"
	"age": <value>		e.g. 21
}
 */

type Student struct {
	Name string `json:"name"`
	Age int `json:"age"`
}


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

func replyWithStudent(w http.ResponseWriter, db *StudentsDB, i int)  {
	if db.Count() < i {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(db.Get(i))
}

func handlerStudent(w http.ResponseWriter, r *http.Request)  {
	//
	db := StudentsDB{}
	if r.Method == "POST"{
		http.Error(w, "not implemented yet", http.StatusNotImplemented)
		return
	}
	//
	http.Header.Add(w.Header(),"content-type", "application/json")
	parts := strings.Split(r.URL.Path, "/")

	if len(parts) == 3{
		if parts[2] == ""{
			replyWithAllStudents(w, &db)
		}else {
			i, err :=strconv.Atoi(parts[2])
			if err != nil{
				http.Error(w,http.StatusText(http.StatusBadRequest) ,http.StatusBadRequest)
				return
			}
			replyWithStudent(w, &db, i)
		}
	}else{
		// handle error
		status := 400
		http.Error(w, http.StatusText(status), status)
		return
	}
}

func main() {
	http.HandleFunc("/hello/", handlerHello)
	http.HandleFunc("/student/", handlerStudent)
	http.ListenAndServe(":8000", nil)
}
