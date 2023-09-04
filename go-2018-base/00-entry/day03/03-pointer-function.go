package main

import "fmt"

func MySwap(a *int, b *int) {
	//c:=*a
	//*a=*b
	//*b=c
	*a, *b = *b, *a
}

func main() {
	a, b := 10, 20
	fmt.Printf("a=%d; b=%d\n", a, b)
	MySwap(&a, &b)
	fmt.Printf("a=%d; b=%d\n", a, b)

	fmt.Println("指针类型作为函数参数")
}
