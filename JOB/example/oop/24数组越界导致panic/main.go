package main

import "fmt"

func main() {
	teata()
	teatb(20)
	teatc()
}

func teata() {
	fmt.Println("aaaaaaaaaaaaaaaaaaaa")
}
func teatb(x int) {
	var a [10]int
	a[x] = 111 // 当选为20
}
func teatc() {
	fmt.Println("cccccccccccccccccccc")
}
