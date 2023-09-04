package main

import "fmt"

//匿名函数
//闭包：一个函数捕获了和它在同一作用域的其他常量和变量

func main() {
	a := 10
	str := "abc"
	//匿名函数
	f1 := func(aa int) {
		fmt.Println("匿名函数1")
		fmt.Println("aa=", aa)
		//匿名函数可以调用外部的变量
		fmt.Printf("a=%d; str=%s\n", a, str)
	}
	f1(20)

	//定义匿名函数的同时调用
	func() {
		fmt.Println("匿名函数2")
		//闭包可以修改外部变量的值
		a = 100
		str = "bbb"
		fmt.Printf("a=%d; str=%s\n", a, str)
	}()
	fmt.Printf("a=%d; str=%s\n", a, str)
	fmt.Println("05 anonymous function")
}
