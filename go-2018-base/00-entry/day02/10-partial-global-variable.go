package main

import "fmt"

//定义在函数外部的变量就是全局变量
var global10 int = 20

func main() {
	//作用域
	{
		a := 10
		fmt.Println("a=", a)
	}
	{
		a := 20
		fmt.Println("a=", a)
	}
	//fmt.Println("a=",a)
	fmt.Println("global: ", global10)
	fmt.Println("10 局部变量")
}
