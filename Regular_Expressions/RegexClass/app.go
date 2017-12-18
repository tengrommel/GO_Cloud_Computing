package main

import (
	"regexp"
	"fmt"
)

func main() {
	pattern := "^[[:alnum:]]"
	input := "fdafd"
	res, err := regexp.MatchString(pattern, input)
	if err != nil{
		fmt.Println(err.Error())
	}
	fmt.Println(res)
}
