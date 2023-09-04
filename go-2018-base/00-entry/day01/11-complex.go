package main

import "fmt"

func main() {
	fmt.Println("11")

	//复数
	var t complex128
	t = 23.23 + 332i
	fmt.Println("t=", t)

	//复数 自动推导类型
	t2 := 3.3 + 5.55i
	fmt.Printf("t2 type:%T\n", t2)

	//通过内建函数 去实部 虚部
	fmt.Println("real(t2)= ", real(t2), " imag(t2)=", imag(t2))

}
