package main

import (
	"regexp"
	"fmt"
)

func main() {
	buf := "abc azc a7c aac 888 a9c tac"
	// (1)解释规则，它会解析正则表达式，如果成功返回一个解释器
	reg1 := regexp.MustCompile(`a\dc`)
	if reg1 == nil{ // 解释失败，返回nil
		fmt.Println("err = ", reg1)
		return
	}
	// 根据规则提取关键信息
	result := reg1.FindAllStringSubmatch(buf, 1)
	fmt.Println("result1 = ", result)
}
