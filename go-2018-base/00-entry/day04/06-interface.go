package main

import "fmt"

//定义接口类型
type Human interface {
	//方法声明，只有声明，没有实现，没有成员变量
	sayHello()
}

type Student5 struct {
	id int
}

type Teacher5 struct {
	group string
}

func (s *Student5) sayHello() {
	fmt.Println("student say hello")
}

func (t *Teacher5) sayHello() {
	fmt.Println("teacher say hello")
}

//多态
func whoSayHello(i Human) {
	i.sayHello()
}

func main() {

	//	接⼝命名习惯以 er 结尾
	//	接口只有方法声明，没有实现，没有数据字段

	var i Human
	s1 := &Student5{111}
	i = s1
	i.sayHello()

	t1 := &Teacher5{group: "中级教师"}
	i = t1
	i.sayHello()

	whoSayHello(s1)
	whoSayHello(t1)

	fmt.Println("06 接口")
}
