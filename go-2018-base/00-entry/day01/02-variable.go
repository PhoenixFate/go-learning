package main

//导入包 必须使用，否则报错
import "fmt"

func main() {

	//1. 声明 var 变量名 类型；
	//变量声明了必须使用，否则报错
	//2. 只是声明没有初始化的变量，默认0
	var a int

	fmt.Println("a = ", a)

	//4.声明多个变量
	var b, c int
	c = 20
	fmt.Printf("b= %d, c= %d\n", b, c)
	fmt.Printf("xxx\n")

	var (
		dd int
		ee float32
	)
	fmt.Printf("dd=%d, ee=%f\n", dd, ee)

	//5.变量初始化
	var d int = 30
	fmt.Println(d)
	var g = 50.2
	fmt.Println(g)

	//6.自动推导类型，必须初始化   !!!!!!!!!!!!!!!!!!!!!!!!
	e := 50
	//%T 变量所属的类型
	fmt.Printf("c type is %T\n", e)

}
