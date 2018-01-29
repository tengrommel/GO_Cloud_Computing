package main

import (
	"sort"
	"fmt"
)

func main() {
	a := []int{3,6,2,12,9,7}
	sort.Ints(a)
	for i, v := range a{
		fmt.Println(i, v)
	}
}
