package main

import "fmt"

func main() {
	//分组密码模式
	//1.按位异或 XOR : exclusive or      bit位 相同为0 不同为1
	//2.ECB electronic code book 电子密码本模式
	//3.CBC cipher block chaining 密码块链模式
	//4.CFB cipher feedback 密码反馈模式
	//5.OFB output feedback 输出反馈模式
	//6.CTR counter 计数器模式

	// ctr cfb ofc 不需要对明文进行填充
	//cbc  ecb    需要对明文进行填充

	var a int = 10
	var b int = 20
	c := a ^ b ^ b
	fmt.Println(c)

	fmt.Println("04 分组密码模式")
}
