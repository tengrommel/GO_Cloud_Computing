package main

import "fmt"

type Person struct {
	name string
	sex byte
	age int
}

// Person 类型，实现了一个方法
func (p *Person)PrintInfo() {
	fmt.Printf("name = %s, sex = %c, age = %d\n", p.name, p.sex, p.age)
}

//有一个学生，继承Person字段，成员方法都继承了
type Student struct {
	Person // 匿名字段
	id int
	addr string
}

// Student也实现了一个方法，这个方法和Person方法同名，这种方法叫重写
func (s *Student)PrintInfo() {
	fmt.Println("Student: tmp = ", s)
}


func main() {
	s := Student{
		Person{"mike", 'm', 18},
		666,
		"bj",
	}
	// 就近原则，先找本作用域的方法，再找继承的方法
	s.PrintInfo()
	// 显示调用继承的方法
	s.Person.PrintInfo()

}