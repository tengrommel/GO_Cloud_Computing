package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type CrewMember struct {
		ID int `json:"id"`
		Name string `json:"name"`
		SecurityClearance int `json:"security_clearance"`
		AcceessCodes []string `json:"acceess_codes"`
	}

	cm := CrewMember{ Name:"Jaro", SecurityClearance:10, AcceessCodes:[]string{"abc","def"}}
	b, err := json.Marshal(&cm)
	if err !=nil{
		fmt.Println("error:", err)
		return
	}
	fmt.Println(string(b))
}
