package main

import (
	"time"
	"fmt"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	i := 0
	for {
		<- ticker.C
		i++
		fmt.Println("i = ", i)
	}

}
