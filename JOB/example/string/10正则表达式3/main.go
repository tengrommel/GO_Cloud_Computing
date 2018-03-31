package main

import (
	"regexp"
	"fmt"
)

func main() {
	// ``  原生字符串
	buf := `
	<div>测试</div>
	<divdfa></div>
	<div>测试1</div>
	<div>测试2</div>
`
	//reg := regexp.MustCompile(`<div>(.*)</div>`)
	reg := regexp.MustCompile(`<div>(?s:(.*))</div>`)
	if reg == nil{
		fmt.Println("Compile error")
		return
	}

	// 提取关键信息
	result := reg.FindAllStringSubmatch(buf, -1)
	fmt.Println(result)
	// 过滤<></>
	for _, text := range result {
		fmt.Println("text[0] = ", text[0]) // 带<></>
	}

}
