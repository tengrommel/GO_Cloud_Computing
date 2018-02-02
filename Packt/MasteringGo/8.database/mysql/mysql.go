package main

import (
	"database/sql"
	"log"
	"fmt"
	"strings"
	_ "github.com/go-sql-driver/mysql"
)

type crewMember struct {
	id								int
	name							string
	secClearance			int
	position					string
}

type Crew []crewMember

func main() {
	db, err := sql.Open("mysql", "gouser:gouser@/Hydra")
	if err != nil {
		log.Fatal("Could not connect, error ", err.Error())
	}
	defer db.Close()

	cw := GetCrewByPositions(db, []string{"'Mechanic'", "'Biologist'"})
	fmt.Println(cw)

	fmt.Println(GetCrewMemberById(db, 11))

	AddCrewMember(db, crewMember{name: "Kaya Gal", secClearance: 4, position: "Biologist"})

	fmt.Println(GetCrewMemberByPosition(db, "Chemist"))

	cr := Crew{
		crewMember{name: "Adam stler", secClearance: 4, position: "Chemist"},
		crewMember{name: "Zach Garph", secClearance: 5, position: "Biologist"},
	}
	CreateCrewMembersByTransaction(db, cr)
}

func CreateCrewMembersByTransaction(db *sql.DB, crews Crew) {
	tx, err := db.Begin()
	if err !=nil{
		log.Fatal("Could not begin tx", err)
	}
	stmt, err := tx.Prepare("INSERT INTO Personnel (Name,SecurityClearance,Position) VALUES (?,?,?);")
	if err != nil{
		tx.Rollback()
		log.Fatal("Could not do select statement ", err)
	}
	defer stmt.Close()

	for _, person := range crews{
		_, err := stmt.Exec(person.name, person.secClearance, person.position)
		if err != nil{
			tx.Rollback()
			log.Fatal("Could not query positions ", err)
		}
	}
	tx.Commit()
	return
}

func AddCrewMember(db *sql.DB, member crewMember) int64 {
	res, err := db.Exec("INSERT INTO Personnel (Name,SecurityClearance,Position) VALUES (?,?,?)", member.name, member.secClearance, member.position)
	if err != nil{
		log.Fatal(err)
	}
	ra,_ := res.RowsAffected()
	id,_ := res.LastInsertId()

	log.Println("Rows Affected", ra, "Last inserted id", id)
	return id
}
func GetCrewMemberById(db *sql.DB, id int) (cm crewMember) {
	row := db.QueryRow("SELECT * FROM Personnel where id = ?", id)
	err := row.Scan(&cm.id, &cm.name, &cm.secClearance, &cm.position)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func GetCrewMemberByPosition(db *sql.DB, position string) (cr Crew) {
	stmt, err := db.Prepare("SELECT * FROM Personnel where Position = ?")
	if err !=nil{
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(position)
	if err !=nil{
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next(){
		cm := crewMember{}
		err = rows.Scan(&cm.id, &cm.name, &cm.secClearance, &cm.position)
		if err != nil{
			log.Fatal(err)
		}
		cr = append(cr, cm)
	}

	if err := rows.Err(); err !=nil{
		log.Fatal(err)
	}
	return
}

func GetCrewByPositions(db *sql.DB, positions []string) Crew {
	Qs := fmt.Sprintf("SELECT id, Name, SecurityClearance, Position from Personnel" +
		"where Position in (%s);",strings.Join(positions, ","))
	rows, err := db.Query(Qs)
	if err != nil{
		log.Fatal("Could not get data from the Personnel table ", err)
	}
	defer rows.Close()

	retVal := Crew{}
	cols, _ := rows.Columns()
	fmt.Println("Columns detected: ", cols)

	for rows.Next(){
		member := crewMember{}
		err = rows.Scan(&member.id, &member.name, &member.secClearance, &member.position)
		if err !=nil{
			log.Fatal("Error scanning row", err)
		}
		retVal = append(retVal, member)
	}

	if err := rows.Err(); err !=nil{
		log.Fatal(err)
	}
	return retVal
}


