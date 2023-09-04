package main

import (
	"fmt"
)

type Student struct {
	id      int
	name    string
	sex     byte
	age     int
	address string
}

func main() {
	//顺序初始化，每个成员必须初始化
	var s1 Student = Student{1, "tom", 0, 22, "中国大陆"}
	fmt.Println(s1)

	//指定成员初始化, 没有初始化的，默认为0
	s2 := Student{name: "mike", sex: 1, age: 88}
	fmt.Println(s2)

	//结构体指针变量
	var s3 = &Student{3, "tom3", 0, 22, "中国大陆3"}
	fmt.Println(*s3)

	s4 := &Student{name: "王大陆", age: 66}
	fmt.Println(*s4)
	fmt.Println(s4.name)
	fmt.Println((*s4).name)
	fmt.Printf("type of s4: %T\n", s4)

	fmt.Println("16 结构体")
}
