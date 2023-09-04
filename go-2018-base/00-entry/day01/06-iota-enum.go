package main

import "fmt"

func main() {
	fmt.Println("06")

	//1.iota常量自动生成器，每一行，自动加1
	//2.iota给常量赋值使用
	const (
		a = 10
		b = iota
		c = iota
	)
	fmt.Printf("a=%d; b=%d; c=%d\n", a, b, c)

	//3.iota遇到const，重置为0
	const d = iota
	fmt.Printf("d=%d\n", d)

	//4.可以只写一个iota
	const (
		a2 = iota
		b2
		c2
	)
	fmt.Printf("a2=%d; b2=%d; c2=%d\n", a2, b2, c2)

	//5.如果同一行，iota则是一个值
	const (
		i          = iota
		j1, j2, j3 = iota, iota, iota
		k          = iota
	)
	fmt.Printf("i=%d; j1=%d,j2=%d,j3=%d; k=%d\n", i, j1, j2, j3, k)
}
