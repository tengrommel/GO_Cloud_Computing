package main

import "fmt"

type Person struct {
	name string
	sex byte
	age int
}

func (p Person)SetInfoValue()  {
	fmt.Println("SetInfoValue")
}

func (p *Person)SetInfoPointer()  {
	fmt.Println("SetInfoPointer")
}


func main() {
	// 结构体是一个指针变量，它能够调用那些方法，这些方法就是一个集合，简称方法集
	p := &Person{"mike", 'm', 18}
	//p.SetInfoPointer()
	//p.SetInfoValue()
	(*p).SetInfoPointer()
	(*p).SetInfoValue()
	// 内部做的转换，先把指针p，转成*p后再调用
	// (*p).SetInfoValue()
}
