package main

import "fmt"

func main() {
	fmt.Println("07")
	//未初始化 默认为false
	var a bool
	fmt.Println("a=", a)

	var b bool = false
	fmt.Println(b)

	var c = false
	fmt.Println(c)

	d := false
	fmt.Printf("d type: %T\n", d)
	fmt.Println("d= ", d)
	fmt.Printf("bool int:%v\n", d)
}
