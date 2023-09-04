package main

import "fmt"

//go函数可以返回多个值
func test() (int, int, int) {
	return 10, 20, 30
}

func main() {
	fmt.Println("04")

	//多重赋值
	a, b, c := 10, 20, 30
	fmt.Printf("a=%d; b=%d; c=%d\n", a, b, c)

	//交换两个变量
	//传统
	var temp int
	temp = a
	a = b
	b = temp
	fmt.Printf("a=%d; b=%d\n", a, b)

	//go
	i := 10
	j := 20
	i, j = j, i
	fmt.Printf("i=%d; j=%d\n", i, j)

	//_ 匿名变量
	a1, _, a2 := test()
	fmt.Printf("a1=%d; a2=%d\n", a1, a2)

}
