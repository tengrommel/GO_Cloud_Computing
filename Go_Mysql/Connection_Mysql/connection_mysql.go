package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	// I've set up a database on my local machine using idea
	// The database is called test
	db, err := sql.Open("mysql", "root:teng@tcp(127.0.0.1:3306)/test")
	if err != nil{
		panic(err.Error())
	}

	// defer the close till after the main function has finished executing
	defer db.Close()

}
