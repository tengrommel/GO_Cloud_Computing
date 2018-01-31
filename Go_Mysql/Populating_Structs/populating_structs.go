package main

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

/*
Tag... - a very simple struct
 */

type Tag struct {
	ID 		int 			`json:"id"`
	Name 	string 		`json:"name"`
	Age   int 			`json:"age"`
}

func main() {
	// Open up our database connection.
	db, err := sql.Open("mysql", "root:teng@tcp(127.0.0.1:3306)/test")

	// if there is an error opening the connection, handle it
	if err != nil{
		log.Print(err.Error())
	}
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT id, name, age FROM stu")
	if err != nil{
		panic(err.Error())
	}

	for results.Next(){
		var tag Tag
		// for each row, scan the result into our tag composite object
		err = results.Scan(&tag.ID, &tag.Name, &tag.Age)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		log.Printf(tag.Name)
	}

	var one Tag
	err = db.QueryRow("SELECT id, name, age FROM stu where id = ?", 2).Scan(&one.ID, &one.Name, &one.Age)
	log.Println(one.Age)
	log.Println(one.Name)
	log.Println(one.ID)
}
