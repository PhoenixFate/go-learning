package main

import "fmt"

type Person42 struct {
	name string
	age  int
}

//一个指向自定义类型的值的指针，它的方法集由该类型定义的所有方法组成，无论这些方法接受的是一个值还是一个指针。
//也就是Person42 实例对象和指针对象，都能调用这个方法
func (p *Person42) printInfo() {
	fmt.Println("person pointer: ", *p)
}

//也就是Person42 实例对象和指针对象，都能调用这个方法
func (p Person42) printValue() {
	fmt.Println("person value: ", p)
}

func main() {
	var p1 = Person42{"tom4", 88}
	p1.printInfo()
	p1.printValue()

	var p2 *Person42 = &p1
	p2.printInfo()
	p2.printValue()

	fmt.Println("05 方法的继承")
}
