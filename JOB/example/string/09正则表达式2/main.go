package main

import (
	"regexp"
	"fmt"
)

func main() {
	buf := "3.14 567 adgasdf 1.23 7. 8.99 fasf21f 6.66"

	// 解释正则表达式，+匹配前一个字符的一次或多次
	reg := regexp.MustCompile(`\d\.\d+`)
	if reg == nil{
		fmt.Println("MustCompile err")
		return
	}
	// 提取关键信息
	//result := reg.FindAllString(buf, -1)
	result := reg.FindAllStringSubmatch(buf, -1)
	fmt.Println("result = ", result)
}
