package main

import (
	"reflect"
	"fmt"
)

func main() {
	var x1 float32 = 5.7
	InspectIfTypeFloat(x1)
}

func InspectIfTypeFloat(i interface{})  {
	v := reflect.ValueOf(i)
	fmt.Println("type: ", v.Type())
	fmt.Println("Is type is float64?", v.Kind()==reflect.Float64)
	fmt.Println("Float Value:", v.Float())
	fmt.Println("Value:", v)
	x2 := v.Float()
	v2 := reflect.ValueOf(x2)
	fmt.Println("Is type is float64?", v.Kind()==reflect.Float64)
	fmt.Println("Type of v2", v2.Type())

	interfaceValue := v.Interface()
	switch t := interfaceValue.(type) {
	case float32:
		fmt.Println("Original float32 value", t)
	}
}