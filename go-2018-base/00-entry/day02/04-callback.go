package main

import "fmt"

type MyFunctionType4 func(int, int) int

//回调函数，函数有一个参数是函数类型，这个函数就是回调函数
func MyCalculate2(a, b int, funcTest MyFunctionType4) (result int) {
	result = funcTest(a, b)
	return
}

func MyAdd2(a, b int) int {
	return a + b
}

func MyMinus2(a, b int) int {
	return a - b
}

func main() {
	result := MyCalculate2(10, 20, MyAdd2)
	fmt.Println("result: ", result)
	result2 := MyCalculate2(10, 20, MyMinus2)
	fmt.Println("result2: ", result2)

	fmt.Println("04 callback")
}
