package main

import (
	"log"
	"os"
	"encoding/json"
)

type CrewMember struct {
	ID									int 				`json:"id,omitempty"`
	Name  							string			`json:"name"`
	SecurityClearance 	int					`json:"clearance level"`
	AccessCode 					[]string		`json:"access_code"`
}

type ShipInfo struct {
	ShipID 			int
	ShipClass		string
	Captain			CrewMember
}


func main() {
	f, err := os.Create("jfile.json")
	PrintFatalError(err)
	defer f.Close()

	//create some data to encode
	cm := CrewMember{Name:"Jaro", SecurityClearance:10, AccessCode:[]string{"ADA", "LAL"}}
	si := ShipInfo{1,"Fighter", cm}

	err = json.NewEncoder(f).Encode(&si)
	PrintFatalError(err)
}

func PrintFatalError(err error)  {
	if err != nil{
		log.Fatal("Error happened while processing file ", err)
	}
}