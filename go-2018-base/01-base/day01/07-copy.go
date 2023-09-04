package main

import "fmt"

func removeData(data []int, index int) []int {
	copy(data[index:], data[index+1:])
	return data[:len(data)-1]
}

func main() {

	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := data[8:]
	s2 := data[:5]
	//copy（目标切片 源切片）
	copy(s2, s1)
	fmt.Println(s2)
	fmt.Println(s1)
	fmt.Println(data)

	data2 := []int{5, 6, 7, 8, 9}
	//去除7这元素
	fmt.Println(removeData(data2, 2))
}
