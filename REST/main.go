package main

import (
	"net/http"
	"fmt"
	"strings"
)


var db StudentsDB

func handlerHello(w http.ResponseWriter, r *http.Request)  {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 4{
		status := 400
		http.Error(w, http.StatusText(status), status)
		return
	}
	name := parts[2]
	fmt.Fprintln(w, parts)
	fmt.Fprintf(w, "Hello %s %s!\n", name, parts[3])
}




func main() {
	db = StudentsDB{}
	db.Init()
	http.HandleFunc("/hello/", handlerHello)
	http.HandleFunc("/student/", handlerStudent)
	http.ListenAndServe(":8000", nil)
}
