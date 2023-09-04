package main

import "fmt"

//匿名字段
type Person2 struct {
	name string
	sex  byte
	age  int
}

type Student2 struct {
	*Person2 //匿名字段，那么默认student就包含了Person的所有字段
	id       int
	address  string
	name     string //和Person里面重名的name
	int             //基础类型的匿名字段，通过s.int 来访问该字段
}

func main() {
	var s1 = Student2{&Person2{"tom", 1, 34}, 2342342, "北京王府井", "aaa", 1}
	fmt.Println(s1)

	var s2 Student2
	s2.Person2 = new(Person2)
	s2.name = "xxx" //如果s2.Person 未初始化的话：panic: runtime error: invalid memory address or nil pointer dereference
	s2.sex = 1
	s2.age = 19
	s2.address = "上海新天地"
	s2.id = 2343
	fmt.Println(s2)

	fmt.Println("结构体 指针类型的成员变量")
}
