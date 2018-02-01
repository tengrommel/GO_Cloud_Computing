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

	m := map[string]int{"item1":1,"item2":2}
	bm, err := json.Marshal(m)
	if err != nil{
		fmt.Println("error", err)
		return
	}
	fmt.Println(string(bm))

	s:=[]CrewMember{cm, CrewMember{2,"Jim",5,[]string{"TLT","RAR"}}}

	bSlice, err := json.Marshal(&s)
	if err!=nil{
		fmt.Println("error:", err)
	}
	fmt.Println(string(bSlice))

}
