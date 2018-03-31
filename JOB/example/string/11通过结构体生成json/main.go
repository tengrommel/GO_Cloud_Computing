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
	// 定义一个结构体变量，同时初始化
	s := IT{"itcast", []string{"Go", "C++", "Python", "Test"}, true, 666.666}
	// 编码，根据内容生成json文本
	//buf, err := json.Marshal(s)
	buf, err := json.MarshalIndent(s, "", " ") // 格式化编码


	if err != nil{
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("buf = ", string(buf))
}
