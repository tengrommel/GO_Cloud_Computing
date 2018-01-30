package main

import (
	"reflect"
	"fmt"
)

type myFloat float64

func main() {
	var x1 float32 = 5.7
	inspectIfTypeFloat(x1)
	var x3 myFloat = 5
	v := reflect.ValueOf(x3)
	fmt.Println(v.Type())
	// Kind 返回元
	fmt.Println(v.Kind() == reflect.Float64)
}

func inspectIfTypeFloat(i interface{})  {
	v := reflect.ValueOf(i)
	fmt.Println("type: ", v.Type())
	fmt.Println("Is type is float64?", v.Kind()==reflect.Float64)
	fmt.Println("Float Value:", v.Float())
	fmt.Println("Value:", v)
	x2 := v.Float()
	v2 := reflect.ValueOf(x2)
	fmt.Println("Is type is float64?", v.Kind()==reflect.Float64)
	fmt.Println("Type of v2", v2.Type())
}