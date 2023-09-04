package main

import "fmt"
import "regexp"

func main() {

	buf := "abc azc a7c aac 888 a9c tac"
	//1.解释规则 它会解析正则表达式 如果成功则返回解释器
	//reg1:=regexp.MustCompile("a.c")
	reg1 := regexp.MustCompile(`a[0-9]c`)
	//reg1:=regexp.MustCompile(`a\dc`)  //[0-9] 等价于 \d
	if reg1 == nil {
		fmt.Println("mustCompile error")
		fmt.Println(reg1)
		return
	}
	//2.根据规则提取关键信息
	//n 表示匹配次数
	result1 := reg1.FindAllStringSubmatch(buf, -1)
	fmt.Println(result1)

	//.表示任意字符（\n除外）
	//\表示转义 \. 表示.
	//*匹配前一次规则 0次或多次
	//+匹配前一次规则 1次或者多次
	//?              0次或1次
	//{m}            m次
	//{m,n}          m-n次
	//{,n}           0-n次
	//{m,}           m-无限次
	//^    匹配字符串开头   ^abc 以abc开头
	//$    匹配字符串结尾   abc$ 以abc结尾
	buffer2 := "234.24 23.4 6.6 67.4"
	reg2 := regexp.MustCompile(`\d+\.\d+`)
	result2 := reg2.FindAllStringSubmatch(buffer2, -1)
	fmt.Printf("result2 type: %T\n", result2)
	fmt.Println(result2)
	fmt.Printf("result2[0] type: %T\n", result2[0])
	fmt.Println(result2[0])

	html := `<div>abc</div>
		<div>abc
		dd
		</div>
		<div>bcd</div>
		<div>哈桑</div>
	`
	reg3 := regexp.MustCompile(`<div>(.*)</div>`)
	result3 := reg3.FindAllStringSubmatch(html, -1)
	fmt.Println(result3)

	fmt.Println("05 正则表达式")
}
