package main

import "fmt"

func main() {
	//函数 copy 在两个 slice 间复制数据，复制⻓度以 len 小的为准，两个 slice 可指向同⼀底层数组。
	s1 := []int{1, 2}
	s2 := []int{6, 6, 6, 6, 6}
	//copy(destination, src)
	copy(s2, s1)
	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Println("12 copy")
}
