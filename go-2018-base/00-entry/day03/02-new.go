package main

import "fmt"

func main() {
	//new(T) 返回*T
	//我们只需使用new()函数，无需担心其内存的生命周期或怎样将其删除，因为Go语言的内存管理系统会帮我们打理一切。
	var p1 *int
	p1 = new(int) //p1 为*int类型，指向匿名的int变量
	fmt.Println("*p1=", *p1)

	//无需担心内存的生命周期或者怎样将其删除， go有自己的内存管理系统
	p2 := new(int)
	*p2 = 77
	fmt.Println("*p2=", *p2)
	fmt.Println("02 new")
}
