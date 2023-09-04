package main

import "fmt"

func main() {
	var str1 string
	str1 = "abc"
	fmt.Println(str1)

	//自动推导类型
	var str2 = "mike"

	fmt.Printf("str2 type: %T\n", str2)
	fmt.Println(str2)

	//内建函数 len()
	fmt.Println("str2 length: ", len(str2))

	var ch1 = 's'
	//字符串隐藏了\0
	var ch2 = "s"
	fmt.Printf("ch1 type: %T\n", ch1)
	fmt.Printf("ch2 type: %T\n", ch2)

	var ss = "abc"
	fmt.Printf("str[0]=%c; str[1]=%c; str[2]=%c", ss[0], ss[1], ss[2])

}
