package main

import "fmt"

func main() {
	var a int = 10
	fmt.Printf("address of a: %p\n", &a)
	fmt.Printf("address of a: %v\n", &a)

	//保存某个变量的地址，需要指针类型 *int
	var p *int
	p = &a
	fmt.Printf("value of p: %x\n", p)
	fmt.Printf("value of p: %v\n", p)
	*p = 20
	fmt.Printf("*p:%d; a=%d\n", *p, a)

	//指针的基本操作
	//默认值nil，没有NULL常量
	//&，*
	//不支持指针运算-> , 直接使用 .  访问目标成员
	var p2 *int = nil
	p2 = &a
	fmt.Printf("value of p2: %x\n", p2)

	fmt.Println("01 pointer")
}
