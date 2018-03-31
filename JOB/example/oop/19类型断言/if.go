package main

import "fmt"

type Student struct {
	name string
	id int
}


func main() {
	i := make([]interface{}, 3)
	i[0] = 1                    // int
	i[1] = "hello go"           // string
	i[2] = Student{"mike", 666} // student

	// 类型查询，类型断言
	// 第一个返回下标，第二个返回下标对应的值，data分别是i[0],i[1],i[2]
	for index, data := range i {
		// 第一个返回值，第二个返回判断结果的真假
		if value, ok := data.(int); ok {
			fmt.Printf("x[%d] 类型为int， 内容为%d\n", index, value)
		} else if value, ok := data.(string); ok == true {
			fmt.Printf("x[%d] 类型为Student， 内容为%s\n", index, value)
		}
	}

	for index, data :=range i {
		switch value := data.(type) {
		case int:
			fmt.Println(index, value)

		}
	}
	var int_ int
	ss := string(int_)
	fmt.Println(ss)
}

