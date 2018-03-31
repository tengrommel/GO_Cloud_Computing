package main

import "fmt"

// 定义接口类型
type Humaner interface {
	// 方法，只有声明，没有实现，由别的类型实现
	sayhi()
}

type Student struct {
	name string
	id int
}

type Teacher struct {
	addr string
	group string
}

// Student指针类型实现了此方法
func (s *Student)sayhi()  {
	fmt.Printf("Student [%s, %d], sayhi\n", s.name, s.id)
}

// Teacher指针类型实现了此方法
func (t *Teacher)sayhi()  {
	fmt.Printf("Teacher [%s, %d], sayhi\n", t.addr, t.group)
}

type MyStr string
// MyStr指针类型实现了此方法
func (t *MyStr)sayhi()  {
	fmt.Printf("MyStr[%s], sayhi\n", *t)
}

// 定义了一个普通函数，函数的参数为接口类型
// 只有一个函数，可以有不同表现，多态
func WhoSayHi(i Humaner)  {
	i.sayhi()
}

func main() {
	// 定义接口类型的变量
	var i Humaner
	// 只要实现了此接口方法的类型，那么这个类型的变量（接受的类型）就可以给i赋值
	s := &Student{"mike", 666}
	i=s
	i.sayhi()

	t := &Teacher{"bj", "go"}
	i=t
	i.sayhi()

	var str MyStr = "hello mike"
	i = &str
	i.sayhi()

	// 调用同一函数，不同表现
	WhoSayHi(s)
	WhoSayHi(t)
	WhoSayHi(&str)

	// 创建一个切片
	x := make([]Humaner, 3)
	x[0] = s
	x[1] = t
	x[2] = &str

	// 第一个返回下标，第二个返回下标所对应的值
	for _, i := range x {
		i.sayhi()
	}
}
