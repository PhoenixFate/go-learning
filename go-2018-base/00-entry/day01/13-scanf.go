package main

//输入
import "fmt"

func main() {
	var a int
	fmt.Println("请输入a的变量:")

	fmt.Scanf("%d", &a)

	fmt.Println("得到a的值: ", a)

	var b int
	fmt.Println("请输入b的变量:")
	fmt.Scan(&b)
	fmt.Println("得到b的值: ", b)

	fmt.Println("13")
}
