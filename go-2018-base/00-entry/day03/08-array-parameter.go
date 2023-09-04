package main

import "fmt"

//数组作函数参数，是值传递
func modify(a [5]int) {
	a[0] = 99
}

//数组指针
func modify2(a *[5]int) {
	(*a)[0] = 88
}

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	modify(a)
	fmt.Println(a)
	modify2(&a)
	fmt.Println(a)

	fmt.Println("数组作为函数参数")
}
