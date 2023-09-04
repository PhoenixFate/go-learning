package main

import "fmt"

func main() {

	var a = 10
	var b = 30
	var c int
	var d int
	//算术运算符
	a = a + b
	b = a - b
	c = 2 * 3
	// 和C语言一样，整除
	d = 11 / 2
	e := 13 % 3
	var f = 1
	//只有后置的++ 后置的--
	f++
	f--
	fmt.Printf("a=%d; b=%d; c=%d; d=%d; e=%d; f=%d\n", a, b, c, d, e, f)

	//关系运算符
	b1 := a == b
	b2 := a != b
	b3 := a > b
	b4 := a >= b
	b5 := a < b
	b6 := a <= b
	fmt.Printf("b1:%v; b2:%v; b3:%v; b4:%v; b5:%v; b6:%v\n", b1, b2, b3, b4, b5, b6)

	//逻辑运算符
	c1 := b1 && b2
	c2 := b1 || b2
	c3 := !b1
	fmt.Printf("c1=%v; c2=%v; c3=%v\n", c1, c2, c3)

	//位运算符
	d1 := 60 & 13 //按位与
	d2 := 60 | 13 //按位或
	d3 := 60 ^ 13 //异或
	d4 := 4 << 2  //左移 左移n位就是乘以2的n次方
	d5 := 4 >> 2  //右移 右移n位就是除以2的n次方
	fmt.Printf("d1:%v; d2:%v; d3:%v; d4:%v; d5:%v;\n", d1, d2, d3, d4, d5)
}
