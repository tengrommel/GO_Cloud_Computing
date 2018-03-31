package main

import (
	"strconv"
	"fmt"
)

func main() {
	// 转换为字符串后追加到字节数组
	slice := make([]byte, 0, 1024)
	slice = strconv.AppendBool(slice, true)
	// 第二个数要追加的数，第三个为指定10进制方式追加
	slice = strconv.AppendInt(slice, 123, 10)
	slice = strconv.AppendQuote(slice, "abcgohello")
	fmt.Println("slice = ", string(slice))

	// 其他类型转换为字符串
	var str string
	str = strconv.FormatBool(false)
	str = strconv.FormatFloat(3.143432, 'f', -1,64)

	// 整型转字符串
	str = strconv.Itoa(6666)
	boo, _ := strconv.ParseBool("true")
	fmt.Println(boo)
	fmt.Println(str)
}
