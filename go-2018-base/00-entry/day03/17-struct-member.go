package main

import (
	"fmt"
)

/*
//student01只能在本文件件引用，因为首字母小写
type student01 struct {
    Id   int
    Name string //成员首字母大写 相当于public
}

//Student02可以在任意文件引用，因为首字母大写
type Student02 struct {
    Id   int
    name string //成员首字母小写 相当于private
}
*/

type Student2 struct {
	id      int
	name    string
	sex     byte
	age     int
	address string
}

//值传递
func testStruct(s Student2) {
	s.name = "我叫爸爸"
}

func testStruct2(s *Student2) {
	s.name = "我叫爸爸"
}

func main() {
	//p.成员 和(*p).成员 操作是等价的
	var s Student2
	//s=nil //error
	//操作成员，点操作符.
	s.name = "杨幂"
	s.age = 32
	s.address = "北京"
	s.sex = 1
	fmt.Println(s)

	var s2 Student2
	var s3 = &s2
	//s3=nil 只有指针才可以赋为nil
	//s3.name (*s3).name 完全等价
	s3.name = "刘恺威"
	s3.age = 42
	s3.address = "北京"
	fmt.Println(s2)

	//通过new 创建结构体变量
	//s4是一个指针变量
	s4 := new(Student2)
	s4.name = "钟文泽"
	s4.age = 23
	s4.address = "深圳"
	fmt.Println(*s4)
	fmt.Printf("s4 type: %T\n", s4)

	//结构体 比较 赋值
	var s6 = Student2{1, "tom", 1, 22, "北京"}
	var s7 = Student2{1, "tom", 1, 22, "北京"}
	var s8 = Student2{2, "tom", 1, 22, "北京"}
	fmt.Println("s6==s7: ", s6 == s7)
	fmt.Println("s7==s8: ", s7 == s8)

	//同类型的两个结构体变量可以相互赋值
	var s9 Student2
	s9 = s8
	fmt.Println("s9: ", s9)

	//!!!!!!!!!!!!!!结构体作函数参数是值传递
	testStruct(s9)
	fmt.Println(s9)
	testStruct2(&s9)
	fmt.Println(s9)

	//！！！！！！！！！！！！
	//函数、结构体、结构体成员，如果首字母小写，只能在同一个包中可见，
	//首字母大写，其他包可见

	fmt.Println("17 结构体操作成员")
}
