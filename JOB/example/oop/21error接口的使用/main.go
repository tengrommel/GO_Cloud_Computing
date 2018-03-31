package main

import (
	"fmt"
	"errors"
)

func main() {
	err1 := fmt.Errorf("%s", "this is normal err")
	fmt.Println(err1)
	err2 := errors.New("ok")
	fmt.Println("err2", err2)
}
