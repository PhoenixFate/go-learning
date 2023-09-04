package main

import "fmt"

func test61() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := test61()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	fmt.Println("06 闭包")
}
