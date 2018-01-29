package main

import (
	"sync"
	"time"
	"math/rand"
	"fmt"
)

type MapCounter struct {
	m map[int]int
	sync.RWMutex
}

func main() {
	mc := MapCounter{m: make(map[int]int)}
	go runWriter(&mc, 10)
	go runReader(&mc, 10)
	go runReader(&mc, 10)
	time.Sleep(15 *time.Second)
}
func runReader(mapCounter *MapCounter, n int) {
	for i:=0;i<=n;i++  {
		mapCounter.Lock()
		mapCounter.m[i]=i*10
		mapCounter.Unlock()
		time.Sleep(1 *time.Second)
	}
}
func runWriter(mapCounter *MapCounter, n int) {
	for{
		mapCounter.RLock()
		v := mapCounter.m[rand.Intn(n)]
		mapCounter.RUnlock()
		fmt.Println(v)
		time.Sleep(1 *time.Second)
	}
}
