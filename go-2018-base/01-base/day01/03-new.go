package main

import "fmt"

func main() {
	// 使用new关键字在heap上面开辟一块内存空间
	var p = new(int) //默认是类型的默认值
	fmt.Printf("*p: %d\n", *p)
	*p = 100
	fmt.Printf("*p: %d\n", *p)

	var str = new(string) //默认是类型的默认值
	//%q 打印go语言格式的字符串
	fmt.Printf("%q\n", *str)
	*str = "abc"
	fmt.Printf("%q\n", *str)

}
