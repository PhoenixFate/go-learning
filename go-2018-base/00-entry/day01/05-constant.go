package main

import "fmt"

func main() {
	fmt.Println("05")

	//常量
	const a int = 30
	//常量不允许修改
	//a=20

	const b = 10
	fmt.Printf("b type: %T\n", b)
	fmt.Printf("b=%d\n", b)

	const c = 10.2
	fmt.Printf("c type: %T\n", c)

	//多个不同类型的变量的声明
	//传统
	var e1 int
	var e2 float64
	e1, e2 = 10, 32.2
	fmt.Printf("e1: %d; e2: %f\n", e1, e2)

	var (
		f1 int
		f2 float64
	)
	f1, f2 = 23, 23.4
	fmt.Printf("f1: %d; f2: %f\n", f1, f2)

	//自动推导类型
	var (
		g1 = 20
		g2 = 23.2
	)
	fmt.Printf("g1: %d; g2: %f\n", g1, g2)

	//常量的多重赋值
	const u, v = 3, 4
	//常量组
	const (
		i1 int     = 30
		i2 float64 = 56.34
	)
	fmt.Printf("i1:%d; i2:%f\n", i1, i2)

	//常量组
	//自动推导类型
	const (
		j1 = 30
		j2 = 56.34
	)
	fmt.Printf("j1:%d; j2:%f\n", j1, j2)

}
