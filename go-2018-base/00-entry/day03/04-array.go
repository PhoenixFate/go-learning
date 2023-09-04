package main

import (
	"fmt"
)

func main() {
	//数组的元素个数必须是常量
	var id [10]int
	for i := 0; i < len(id); i++ {
		id[i] = i
	}
	for i := 0; i < len(id); i++ {
		fmt.Printf("id[%d]=%d\n", i, id[i])
	}

	for index, value := range id {
		fmt.Printf("id[%d]=%d\n", index, value)
	}
	//内置函数 len(长度) 和 cap(容量) 都返回数组⻓度 (元素数量)：
	fmt.Printf("len:%d \t;cap:%d\n", len(id), cap(id))
	fmt.Printf("type array:%T\n", id)

	const b = 10
	var arr [b]int
	for i := 0; i < len(arr); i++ {
		fmt.Printf("arr[%d]=%d\n", i, arr[i])
	}

	//定义数组的时候初始化
	//全部初始化
	var arr2 [5]int = [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr2); i++ {
		fmt.Printf("arr2[%d]=%d\n", i, arr2[i])
	}

	//部分初始化, 没有初始化的元素，默认为0
	arr3 := [5]int{1, 2, 3}
	for i := 0; i < len(arr3); i++ {
		fmt.Printf("arr3[%d]=%d\n", i, arr3[i])
	}

	bb := [...]int{1, 2, 3} // 通过初始化值确定数组长度
	fmt.Println(bb)
	fmt.Printf("type bb:%T\n", bb) //[3]int

	bc := []int{1, 2, 3}
	fmt.Println(bc)
	fmt.Printf("type bc:%T\n", bc) //[]int

	//通过索引号初始化元素，未初始化元素值为 0
	bd := [5]int{2: 100, 4: 200}
	fmt.Println(bd)

	fmt.Println("04 array")
}
