package main

import "fmt"

type Humaner interface { // 子集
	sayhi()
}

type Personer interface { // 超集
	Humaner // 匿名字段，继承了sayhi()
	sing(lrc string)
}

type Student struct {
	name string
	id int
}

func (s *Student)sayhi()  {
	fmt.Printf("Student [%s, %d] sayhi\n", s.name , s.id)
}

func (s *Student)sing(lrc string)  {
	fmt.Println("Student在唱着:", lrc)
}


func main() {
	// 超集可以转化为子集， 反过来不可以
	var iPro Personer // 超集
	iPro = &Student{"mike", 666}
	var i Humaner // 子集

	i = iPro
	i.sayhi()

	fmt.Println(iPro)
	fmt.Println(i)
}

