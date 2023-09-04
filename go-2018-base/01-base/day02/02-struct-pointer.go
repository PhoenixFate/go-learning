package main

import "fmt"

type Person2 struct {
	name string
	sex  byte
	age  int
}

type Student2 struct {
	p     Person2
	id    int
	score []int
}

func test2(p *Person2) {
	p.name = "test2"
}

func main() {

	//结构体指针
	//1.顺序初始化
	var man *Person2 = &Person2{"andy", 'm', 30}
	fmt.Println("man: ", *man)

	woman := &Person2{"woman", 'f', 30}
	fmt.Println("woman: ", *woman)

	//使用new关键字创建结构体变量
	p := new(Person2)
	p.name = "p"
	p.sex = 'm'
	p.age = 12
	fmt.Printf("p type: %T\n", p)
	fmt.Println("p: ", *p)

	//结构体指针传参
	test2(p)
	fmt.Println("p: ", *p)

}
