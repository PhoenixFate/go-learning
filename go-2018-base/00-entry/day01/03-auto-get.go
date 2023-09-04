package main

//自动类型推导
import "fmt"

func main() {

	var a int
	a = 20
	fmt.Println(a)

	//同一个变量只能使用一次
	b := 20
	fmt.Println(b)

	b = 30

	//一段一段输出，自动加换行
	fmt.Println(b)
	//格式化输出
	fmt.Printf("b= %d\n", b)

	c := 40
	fmt.Println("a=", a, " b=", b, " c=", c)
	fmt.Printf("a=%d b=%d c=%d \n", a, b, c)

	var d = 50
	fmt.Printf("d type: %T\n", d)
	fmt.Printf("d=%d\n", d)
}
