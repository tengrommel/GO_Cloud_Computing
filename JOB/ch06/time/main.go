package main

import (
	"fmt"
	"time"
)

func main() {
	var n time.Duration
	n = time.Hour
	// 更好的打印
	fmt.Println(n.String())
}
