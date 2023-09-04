package main

import "fmt"

//栈帧
/*
	栈帧
	用来给函数提供运行内存空间。取内存于stack上
	当函数调用时，产生栈帧 函数调用结束时，释放栈帧
	栈帧存储：1.局部变量  2.形参 3.内存字段描述值（栈基指针、栈顶指针）
*/
func test() {
	var b int = 1000
	fmt.Printf("b: %d\n", b)
}

func main() {
	var a int = 10
	fmt.Printf("&a 的地址%p\n", &a)
	fmt.Printf("&a 的地址%x\n", &a)
	var p *int = &a
	//*p 称为解引用、间接引用
	fmt.Printf("p 指向的值%d\n", *p)
	fmt.Printf("p 的值%x\n", p)
	fmt.Printf("p 的地址%p\n", &p)
	//var pp *int=&p  //go语言中不支持二级指针
	test()
}
