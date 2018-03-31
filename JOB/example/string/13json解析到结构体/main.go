package main

import (
	"encoding/json"
	"fmt"
)

// 成员变量名首字母必须大写
type IT struct {
	Company string `json:"company"`
	Subjects []string `json:"subjects"`
	IsOk bool `json:"is_ok"`
	Price float64 `json:"price"`
}

func main() {
	jsonBuf := `
{
 "company": "itcast",
 "subjects": [
  "Go",
  "C++",
  "Python",
  "Test"
 ],
 "is_ok": true,
 "price": 666.666
}`
	var tmp IT
	err :=	json.Unmarshal([]byte(jsonBuf), &tmp)
	if err != nil{
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("tmp = ", tmp)
}
