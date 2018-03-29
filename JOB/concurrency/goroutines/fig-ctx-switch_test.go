package main

import (
	"testing"
	"sync"
)

/*
I've used a few things we haven't discussed yet, so if anything is confusing, just follow
the callouts and focus on the result.
 */

func BenchmarkContextSwitch(b *testing.B)  {
	var wg sync.WaitGroup
	begin := make(chan struct{})
	c := make(chan struct{})

	var token struct{}
	sender := func() {
		defer wg.Done()
		<-begin
		for i:=0;i<b.N ;i++  {
			c <- token
		}
	}
	receiver := func() {
		defer wg.Done()
		<- begin
		for i:=0;i<b.N ;i++ {
			<-c
		}
	}
	wg.Add(2)
	go sender()
	go receiver()
	b.StartTimer()
	close(begin)
	wg.Wait()
}
