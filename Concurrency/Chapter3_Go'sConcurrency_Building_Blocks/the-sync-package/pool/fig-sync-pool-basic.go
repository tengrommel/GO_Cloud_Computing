package main

import (
	"sync"
	"fmt"
	"reflect"
)

func main() {
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instances.")
			return struct {}{}
		},
	}
	myPool.Get()
	/*
	Here we call Get on the pool.
	These calls will invoke the New function defined on the pool since instances haven't yet been instantiated.
	 */
	instance := myPool.Get()
	myPool.Put(instance)
	/*
	Here we put an instance previously retrieved back in the pool.
	number of in stances to one,
	 */
	instanceTwo :=myPool.Get()
	/*
	When this call is executed, we will reuse the instance previously allocated and put it back in the pool.
	The New function will not be invoked.
	*/
	fmt.Println(reflect.TypeOf(instanceTwo))
}
