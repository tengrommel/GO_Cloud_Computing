package main

import (
	"fmt"
	"time"
)

func main()  {
	// 代码的可读性更高
	var n time.Duration
	n = 2 * time.Millisecond
	time.Sleep(n)
	fmt.Println(time.Now())
}
