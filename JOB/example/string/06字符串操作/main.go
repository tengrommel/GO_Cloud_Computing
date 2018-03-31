package main

import (
	"fmt"
	"strings"
)

func main() {
	// "hellogo"中是否包含"hello"，包含返回true，不包含返回false
	fmt.Println(strings.Contains("hellogo", "hello"))
	fmt.Println(strings.Contains("hellogo", "abc"))

	// Joins 组合
	s := []string{"abc", "hello", "mike", "go"}
	fmt.Println(strings.Join(s, ","))
	fmt.Println(strings.Index("abcdhello", "hello"))
	fmt.Println(strings.Index("abcdhello", "go"))
	buf := strings.Repeat("go", 3)
	fmt.Println("buf = ", buf)

	buf = "hello@abc@go@mike"
	s2 := strings.Split(buf, "@")
	fmt.Println(s2)
	// Trim 去掉两头的字符
	buf = strings.Trim("    areyou ok      ", " ")
	fmt.Println(buf)
	s3 := strings.Fields("       are you ok ")
	for _, data := range s3 {
		fmt.Println(data)
	}

}
