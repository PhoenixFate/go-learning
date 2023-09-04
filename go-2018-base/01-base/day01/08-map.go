package main

import (
	"fmt"
)

func main() {
	//map的创建
	var m1 map[int]string //类似于map声明，没有空间，此时map为空 不能操作
	//m1[0] = "abc" //panic: assignment to entry in nil map
	fmt.Println(m1)
	if m1 == nil {
		fmt.Println("m1 is nil")
	}

	m2 := map[int]string{}
	fmt.Println("m2 length: ", len(m2))
	m2[4] = "abc"
	m2[100] = "a100"
	m2[101] = "a101"
	fmt.Println("m2 length: ", len(m2))

	fmt.Println(m2)

	m3 := make(map[int]string) //未指定容量，默认len=0
	fmt.Println(len(m3))
	m3[10] = "m3"
	fmt.Println(m3)

	m4 := make(map[int]string, 5) //5是length，不是cap，map没有cap, 不能在map中使用map
	m4[4] = "m4"
	fmt.Println("m4 length: ", len(m4))
	fmt.Println(m4)

	//定义的同时初始化
	//1.
	var mm1 map[int]string = map[int]string{1: "luffy", 2: "tom"}
	fmt.Println(mm1)

	//2.
	mm2 := map[int]string{1: "abc", 2: "bcd"}
	fmt.Println(mm2)

}
