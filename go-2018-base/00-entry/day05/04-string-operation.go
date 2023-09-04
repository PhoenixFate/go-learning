package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//contains
	fmt.Println("contains: ", strings.Contains("seaFood", "sea"))

	//join
	s := []string{"abc", "bcd", "123"}
	//strings.Join 如果连接符为空，即可以把所有数组元素拼接起来
	fmt.Println("join: ", strings.Join(s, "-"))

	//index
	fmt.Println("index: ", strings.Index("chicken", "ken"))

	//repeat
	fmt.Println("repeat: ", strings.Repeat("abc", 2))

	//replace
	//n表示替换的次数，如果n小于0 表示全部替换
	fmt.Println(strings.Replace("hello hello hello", "o", "oog", 2))
	fmt.Println(strings.ReplaceAll("hello hello hello", "o", "oog"))

	//split
	fmt.Println(strings.Split("a-b-c", "-"))

	//trim
	//%q输出的内容带“”
	fmt.Printf("%q\n", strings.Trim("   hello world   ", " "))

	//fields
	//去除s字符串的空格，并按照空格分隔返回slice
	fmt.Println(strings.Fields("  foo is good "))

	//字符串转换
	//append
	str := make([]byte, 0, 100)
	str = strconv.AppendBool(str, true)
	str = strconv.AppendInt(str, 8888, 10) //以10进制的方式追加
	str = strconv.AppendQuote(str, "abc")
	fmt.Println(string(str))

	//其他类型转换为字符串
	//format
	var str1 string
	str1 = strconv.FormatBool(false)
	fmt.Println(str1)
	str1 = strconv.FormatInt(666, 10)
	fmt.Println(str1)
	str1 = strconv.Itoa(777)
	fmt.Println(str1)

	//字符串转其他类型
	//parse
	flag, _ := strconv.ParseBool("true")
	fmt.Println(flag)
	a, _ := strconv.ParseInt("233", 10, 8)
	fmt.Printf("type a: %T\n", a)
	fmt.Println(a)

	num, _ := strconv.Atoi("354")
	fmt.Println(num)

	fmt.Println("04 字符串操作")
}
