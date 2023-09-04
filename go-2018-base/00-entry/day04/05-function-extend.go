package main

import "fmt"

type Person4 struct {
	name string
	age  int
}

func (p *Person4) printInfo() {
	fmt.Println("person: ", *p)
}

func (p Person4) printInfoValue() {
	fmt.Println("person value: ", p)
}

//默认Student4 会继承Person4的方法
type Student4 struct {
	Person4
	id int
}

//方法的重写
func (s *Student4) printInfo() {
	fmt.Println("student: ", *s)
}

func main() {
	var p1 = Person4{"tom4", 88}
	p1.printInfo()

	var s1 = Student4{Person4{"abc", 99}, 123}
	//Student4继承了Person4的方法
	s1.printInfo()
	//显示调用
	s1.Person4.printInfo()

	/**
	类似于我们可以对函数进行赋值和传递一样，方法也可以进行赋值和传递。
	根据调用者不同，方法分为两种表现形式：方法值和方法表达式。
	两者都可像普通函数那样赋值和传参，区别在于方法值绑定实例，⽽方法表达式则须显式传参。
	*/

	//方法值：方法的入口地址，隐藏了接收者 函数指针指向的值
	pFunc := p1.printInfo //这个就是方法值 调用函数时 无需再传递接收者 隐式传递了接收者 ;方法值，隐式传递 receiver
	pFunc()

	//方法表达式：方法的入口地址，但没有隐藏接收者 须显式传参
	pFunc2 := (*Person4).printInfo
	pFunc2(&p1)

	Func3 := Person4.printInfoValue
	Func3(p1)

	fmt.Println("05 方法的继承")
}
