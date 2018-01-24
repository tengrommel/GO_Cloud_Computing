package main

import "fmt"

/*
管道模式
函数名称必须是声明式的
输出可以作为输入中的一个参数
*/

func main() {
	multiply := func(values []int, multiplier int) []int {
		multipliedValues := make([]int, len(values))
		for i, v := range values{
			multipliedValues[i] = v*multiplier
		}
		return multipliedValues
	}

	add := func(values []int, additive int) []int {
		addedValues := make([]int, len(values))
		for i, v := range values{
			addedValues[i] = v+additive
		}
		return addedValues
	}
	ints := []int{1,2,3,4}
	for _, v := range multiply(add(multiply(ints, 2), 1), 2){
		fmt.Println(v)
	}
}
