package main

import (
	"fmt"
)

func main() {
	//%T 输出变量类型
	a := 10
	b := 3.249
	c := 'c'
	d := "sfs"
	e := true
	fmt.Printf("a:%T; b:%T; c:%T; d:%T\n", a, b, c, d)
	//%d %f %s %c %t(输出bool类型)
	fmt.Printf("a=%d; b=%f; c=%c; d=%s; e=%t\n", a, b, c, d, e)
	//%v 自动匹配格式
	fmt.Printf("a=%v; b=%v; c=%v; d=%v\n", a, b, c, d)
	fmt.Println("12")
}
