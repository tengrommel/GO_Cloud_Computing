package main

import (
	"runtime"
	"fmt"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4) //?
	result := 0
	ch := make(chan int)

	num := 12344/4
	fmt.Println(num)

	for j:=0;j<4;j++ {
		go func(j int) {
			var res int = 0
			for i:=j*num;i<(j+1)*num;i++ {
				if i%2 != 0{
					res = res + i
				}
			}
			fmt.Println(res)
			ch <- res
		}(j)
	}

	go func() {
		for n:= range ch{
			result = result+n
		}
	}()

	time.Sleep(2*time.Second)
	fmt.Println(result+ 12345)
}

