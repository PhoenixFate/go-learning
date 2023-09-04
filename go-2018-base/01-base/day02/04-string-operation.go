package main

import (
	"fmt"
	"strings"
)

func main() {

	//常用的字符串操作
	str := "I love my work and I love my family too"
	//按照指定字符 拆分
	split := strings.Split(str, " ")
	fmt.Println(split)

	//按照空格拆分
	split2 := strings.Fields(str)
	fmt.Println(split2)

	//判断字符串后缀
	flag1 := strings.HasSuffix(str, "too")
	fmt.Println(flag1)
	//判断字符串前缀
	flag2 := strings.HasPrefix(str, "I")
	fmt.Println(flag2)
}
