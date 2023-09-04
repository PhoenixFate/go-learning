package main

import "fmt"

func test71(a int) {
	fmt.Println(1 / a)
}

func main() {

	//多个defer是后进先出
	//defer 延迟调用，main函数结束前调用
	defer fmt.Println("a")
	defer fmt.Println("b")

	//如果调用的函数内部发送错误，其他defer依然会执行
	defer test71(0)

	defer fmt.Println("c")

	fmt.Println("07 defer")
}
