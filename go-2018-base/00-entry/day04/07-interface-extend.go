package main

import "fmt"

type Human7 interface {
	sayHello()
}

//如果一个interface1作为interface2的一个嵌入字段，那么interface2隐式的包含了interface1里面的方法。
type Person7 interface {
	Human7 ////这里像写了SayHi()一样
	sing()
}

type Student7 struct {
	id int
}

func (s *Student7) sayHello() {
	fmt.Printf("student(id=%d) says hello\n", s.id)
}

func (s *Student7) sing() {
	fmt.Printf("student(id=%d) sings\n", s.id)
}

func main() {
	var p Person7
	var s1 = Student7{77}
	p = &s1
	p.sayHello()
	p.sing()

	//子集可以转换为父集，但反过来不行
	var pHuman Human7
	var s2 = Student7{88}
	pHuman = &s2
	pHuman.sayHello()

	//子集可以转换为父集，但反过来不行
	pHuman = p
	pHuman.sayHello()

	fmt.Println("07 接口之间的继承")
}
