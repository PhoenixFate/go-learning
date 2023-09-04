package main

import "fmt"

//切片并不是值传递，而是引用传递，但数组是值传递
func test(a []int) {
	a[0] = 99
}

func main() {
	s1 := []int{1, 2, 5, 6}
	test(s1)
	fmt.Println(s1)

	fmt.Println("13 切片作为函数参数")
}
