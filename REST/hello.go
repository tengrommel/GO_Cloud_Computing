package main

import (
	"net/http"
	"fmt"
	"strings"
	//"encoding/json"
	"encoding/json"
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

func handlerStudent(w http.ResponseWriter, r *http.Request)  {
	http.Header.Add(w.Header(),"content-type", "application/json")
	parts := strings.Split(r.URL.Path, "/")
	s0 := Student{"Tom", 21}
	s1 := Student{"Mark", 16}
	students := []Student{
		s0,
		s1,
	}
	if len(parts) == 3{
		if parts[2] == ""{
			json.NewEncoder(w).Encode(students)
		}else {
			if parts[2] == "0"{
				json.NewEncoder(w).Encode(s0)
			}else if parts[2] == "1"{
				json.NewEncoder(w).Encode(s1)
			}
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
