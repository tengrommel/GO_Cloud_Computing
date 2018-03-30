package main

import (
	"fmt"
	. "os"
) // 导入包，必须使用，否则编译不过


func main() {
	fmt.Println("This is a test")
	fmt.Println("os.Args = ", Args)
}
