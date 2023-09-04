package main

import "fmt"

func main() {
	//数组和切片定义的区别：
	//数组 [] 内需要指定数组长度
	//创建切片 []内需要为空 或者[...]

	//切片的本质，不是数组指针，是一种数据结构，用来操作数组内部元素

	array := [6]int{10, 20, 30, 50, 0, 0}
	fmt.Printf("array address:%p\n", &array)
	//切片名称[low:high:max] [low high)
	//low 起始下标位置
	//high 结束下标位置 取不到
	//len= high-row
	//容量 cap=max-low
	slice := array[1:3:6]
	fmt.Println("slice: ", slice)
	fmt.Printf("slice address: %p\n", &slice)
	fmt.Println("len s: ", len(slice))
	fmt.Println("cap s: ", cap(slice))

	//如果定义切片未指定max，则max为原数组的length
	//cap为length-low
	s2 := array[0:3]
	fmt.Println("s2: ", s2)
	fmt.Println("len s2: ", len(s2))
	fmt.Println("cap s2: ", cap(s2))

	//在切片的基础上截取
	s3 := slice[1:4:5]
	fmt.Println("s3: ", s3)
	fmt.Println("len s3: ", len(s3))
	fmt.Println("cap s3: ", cap(s3))

	//切片的创建
	//1.自动推导
	ss := []int{1, 2, 3, 5}
	fmt.Println(ss)

	//2.make(类型，长度，容量）
	ss2 := make([]int, 10, 20)
	fmt.Println(ss2)
	fmt.Println(cap(ss2))

	//3.make(类型, 长度) 创建切片时候，没有指定容量，容量==长度
	ss3 := make([]int, 5)
	fmt.Println(ss3)
	fmt.Println(cap(ss3))

	//make 只能创建slice map channel，并返回一个有初始值的对象
	//切片做函数参数 传递的是引用
}
