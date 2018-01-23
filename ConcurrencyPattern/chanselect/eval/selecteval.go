package main

import "fmt"
/*
在开始执行select语句的时候，所有跟在case关键字右边的发送语句中的通道表达式都会先求值，
无论它们所在的case是否可能被选择都会是这样
 */

var intChan1 chan int
var intChan2 chan int
var channels = []chan int{intChan1, intChan2}

var numbers = []int{1,2,3,4,5}

func main() {
	select {
	case getChan(0) <- getNumber(0):
		fmt.Println("The 1th case is selected.")
	case getChan(1) <- getNumber(1):
		fmt.Println("The 1th case is selected.")
	default:
		fmt.Println("Default!")
	}
}

func getNumber(i int) int {
	fmt.Printf("numbers[%d]\n", i)
	return numbers[i]
}

func getChan(i int) chan int {
	fmt.Printf("channels[%d]\n", i)
	return channels[i]
}