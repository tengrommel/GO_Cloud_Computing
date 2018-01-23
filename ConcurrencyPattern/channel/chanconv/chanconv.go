package main

import "fmt"
/*
都不能装换
 */

func main() {
	var ok bool
	ch := make(chan int, 1)

	_, ok = interface{}(ch).(<-chan int)
	fmt.Println("chan int => <-chan int:", ok)
	_, ok = interface{}(ch).(chan<- int)
	fmt.Println("chan int => chan<- int:", ok)

	sch := make(chan<- int, 1)
	_, ok = interface{}(sch).(chan int)
	fmt.Println("chan<- int => chan int:", ok)

	rch := make(<-chan int, 1)
	_, ok = interface{}(rch).(chan int)
	fmt.Println("<-chan int => chan int:", ok)

}
