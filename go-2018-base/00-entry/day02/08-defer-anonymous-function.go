package main

import "fmt"

func main() {
	a := 10
	b := 20
	defer func(a, b int) {
		fmt.Println("a=", a, "; b=", b)
	}(a, b) //这里先传参数过去，所以里面是ab是原来的值； 已经先传递参数，只是没有调用
	a = 111
	b = 222

}

func main81() {
	a := 10
	b := 20
	defer func() {
		fmt.Println("a=", a, "; b=", b)
	}()
	a = 111
	b = 222

	fmt.Println("08 defer 和匿名函数")
}
