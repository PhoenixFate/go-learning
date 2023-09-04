package main

import (
	"fmt"
	"unsafe"
)

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	p     Person
	id    int
	score []int
}

func test(man Person) {
	man.name = "test"
	man.age = 23
	man.sex = 'f'
}

func main() {
	//1.顺序初始化：依次将结构体内部所有成员初始化
	var man Person = Person{"andy", 'm', 20}
	fmt.Println("man: ", man)
	fmt.Printf("man: %v\n", man)

	//2.指定成员初始化
	woman := Person{name: "rose", age: 29}
	fmt.Printf("woman:%v\n", woman)

	//索引成员变量 "."
	fmt.Printf("woman.name:%q, age=%d\n", woman.name, woman.age)

	//3.先定义 后初始化
	var superman Person
	superman.name = "superman"
	superman.age = 88
	superman.sex = 'm'
	fmt.Println("superman: ", superman)

	//结构体之间的比较
	//结构体之间比较只能 == !=
	var p1 Person = Person{"p", 'm', 18}
	p2 := Person{"p", 'm', 18}
	var p3 Person = Person{"p3", 'm', 18}
	fmt.Println("p1==p2?: ", p1 == p2)
	fmt.Println("p1==p3?: ", p1 == p3)

	//相同类型结构体的赋值
	//未初始化 则各个成员变量都是类型的默认值
	var temp Person
	fmt.Println("temp: ", temp)
	temp = p3
	fmt.Println("temp: ", temp)
	//获得变量的内存大小
	fmt.Println("sizeof temp: ", unsafe.Sizeof(temp))

	//结构体传参数
	//函数参数传递结构体变量 值传递 将实参的值拷贝一份给形参
	//结构体值传递几乎不用
	test(temp)
	fmt.Println("temp: ", temp)

	//结构体的地址 和 结构体的第一个成员的地址是同一个
	fmt.Printf("temp address: %p\n", &temp)
	fmt.Printf("temp name address: %p\n", &(temp.name))
}
