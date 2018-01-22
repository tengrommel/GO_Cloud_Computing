package main

import (
	"fmt"
)

/*
Ad hoc confinement is when you achieve confinement through a convention - weather
it be set by the languages community, the group you work within, or the codebase you work within.
*/

func main() {
	data := make([]int, 4)
	loopData := func(handleData chan<- int) {
		defer close(handleData)
		for i:= range data{
			handleData <- data[i]
		}
	}
/*
We can see that the data slice of integers is available from both the loopData function and the loop over the handleData channel.
 */
	handleData := make(chan int)
	go loopData(handleData)

	for num := range handleData{
		fmt.Println(num)
	}
}

/*
Set up this way, it is impossible to utilize the channels in this small example. This is a good lead-in to
confinement, but probably not a very interesting example since channels are concurrent-safe.
 */