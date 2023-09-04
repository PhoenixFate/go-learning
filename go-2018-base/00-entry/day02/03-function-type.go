package main

import "fmt"

func MyAdd(a int, b int) int {
	return a + b
}

//不支持函数同名但 参数不同
func MyAddMore(a int, b int, c int) int {
	return a + b + c
}

func MyMinus(a int, b int) int {
	return a - b
}

//MyFunction是一个函数类型
type MyFunction func(int, int) int

func main() {

	//声明一个函数类型的变量
	var funcTest MyFunction
	funcTest = MyAdd
	result := funcTest(10, 20)
	fmt.Println("result: ", result)

	fmt.Println("03 function type")
}
