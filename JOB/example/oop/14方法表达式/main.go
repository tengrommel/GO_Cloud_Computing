package main

import "fmt"

type Person struct {
	name string
	sex byte
	age int
}

// Person 类型，实现了一个方法
func (p Person)SetInfoValue() {
	fmt.Printf("SetInfoValue: %p, %v\n", &p, p)
}
func (p *Person)SetInfoPointer()  {
	fmt.Printf("SetInfoPointer: %p, %v\n", p, p)
}

func main() {
	p := Person{"mike", 'm', 18}
	fmt.Printf("main: %p, %v\n", &p, p)

	// 方法值 f := p.SetInfoPointer
	f := (*Person).SetInfoPointer
	f(&p)
	f2 := (Person).SetInfoValue
	f2(p)
}
