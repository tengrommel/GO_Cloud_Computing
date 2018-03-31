package main

import "fmt"

func main() {
	teata()
	teatb()
	teatc()
}

func teata() {
	fmt.Println("aaaaaaaaaaaaaaaaaaaa")
}
func teatb() {
	//fmt.Println("bbbbbbbbbbbbbbbbbbbb")
	panic("this is a panic test")
}
func teatc() {
	fmt.Println("cccccccccccccccccccc")
}
