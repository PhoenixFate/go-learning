package main

import "fmt"

func main() {
	//range用于迭代遍历
	str := "abc"
	//range 返回两个参数，第一个是元素的位置，一个是元素本身
	for position, char := range str {
		fmt.Printf("str[%d]=%c\n", position, char)
	}

	arr := []int{12, 22, 32, 42}
	fmt.Println(arr)
	for _, data := range arr {
		fmt.Println(data)
	}

	fmt.Println("20 range")
}
