package main

import (
	"sync"
	"fmt"
)

type safeCounter struct {
	i int
	sync.Mutex
}

func (safeCounter *safeCounter) Increment() {
	safeCounter.Lock()
	safeCounter.i++
	safeCounter.Unlock()
}
func (safeCounter *safeCounter) Decrement() {
	safeCounter.Lock()
	safeCounter.i--
	safeCounter.Unlock()
}
func (safeCounter *safeCounter) GetValue() int {
	safeCounter.Lock()
	v := safeCounter.i
	safeCounter.Unlock()
	return v
}

func main() {
	sc := new(safeCounter)
	for i:=0;i<100;i++{
		go sc.Increment()
		go sc.Decrement()
	}
	fmt.Println(sc.GetValue())
}


