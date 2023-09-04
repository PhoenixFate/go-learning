package main

import "fmt"

//类型别名
func main() {

	type bigint int64
	var a bigint
	a = 20
	fmt.Printf("a type:%T\n", a)
	fmt.Printf("a=%d\n", a)

	type (
		myInt int32
		myStr string
	)

	var b myInt
	var c myStr
	b = 30
	c = "abc"
	fmt.Printf("b type:%T\n", b)
	fmt.Println(b)
	fmt.Printf("c type:%T\n", c)
	fmt.Println(c)

	fmt.Println("15")
}
