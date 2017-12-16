package main

import "fmt"

func f() func(int, int) int {
	return func(a int, b int) int {
		return a + b
	}
}

func main() {
	a := f()
	fmt.Println(a(10, 20))
}
