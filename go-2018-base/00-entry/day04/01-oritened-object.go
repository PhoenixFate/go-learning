package main

import "fmt"

//匿名字段
type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	Person  //匿名字段，那么默认student就包含了Person的所有字段
	id      int
	address string
	name    string //和Person里面重名的name
	int            //基础类型的匿名字段，通过s.int 来访问该字段
}

func main() {
	var s1 = Student{Person{"tom", 1, 34}, 2342342, "北京王府井", "aaa", 1}
	fmt.Println(s1)
	//就近原则
	fmt.Println(s1.name)
	//显示调用
	fmt.Println(s1.Person.name)
	//自动推导类型
	s2 := Student{Person{"tom2", 1, 44}, 2342343, "北京王府井2", "bbb", 2}
	//默认就近原则，操作 是Student里面的name
	s2.name = "我是你爸爸"
	//%+v 可以以更详细的格式输出内容
	fmt.Printf("%+v\n", s2)
	//指定成员初始化
	s3 := Student{id: 9999}
	s3.Person.name = "person name"
	s3.Person = Person{"微软", 2, 239}
	fmt.Printf("%+v\n", s3)

	fmt.Println("01 面向对象 入门")
}
