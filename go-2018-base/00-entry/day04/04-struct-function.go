package main

import "fmt"

type Person3 struct {
	name string
	age  int
}

//带有接受者的函数叫方法
//结构体是值传递
func (p Person3) PrintInfo() {
	p.name = "temp"
	fmt.Println(p)
}

//不支持函数重载
//func (p Person3)PrintInfo(a int){
//	p.name="temp"
//	fmt.Println(p)
//}

type PP *Person3

//接受者的基类型不能是指针
//func (p PP)PrintInfo2(){
//	p.name="temp"
//	fmt.Println(p)
//}

//通过函数，给结构体初始化
func (p *Person3) setInfo(name string, age int) {
	p.name = name
	p.age = age
}

func main() {
	p := Person3{"tom", 99}
	p.PrintInfo()
	fmt.Println(p)

	p2 := &Person3{}
	p2.setInfo("老虎", 88)
	fmt.Println(*p2)

	fmt.Println("04 为结构体绑定方法")
}
