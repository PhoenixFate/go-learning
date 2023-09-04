package main

import (
	"bytes"
	"fmt"
)

//填充函数
//如果最后一个分组的字节数不够，填充
//如果字节数刚刚好，则添加一个新分组
//填充个数的值==缺少的字节数
//plainText 明文
func paddingLastGroup(plainText []byte, blockSize int) []byte {
	//1.求出最后一个组中剩余的字节数
	lastNum := len(plainText) % blockSize
	padNum := blockSize - lastNum
	//2.创建新的切片  长度==padNum 每个字节值 byte(padNum)
	padChar := []byte{byte(padNum)} //字节切片，长度1
	//创建需要追加的切片，并且初始化
	addPlain := bytes.Repeat(padChar, padNum)
	//3.将需要追加的串，追加到原始明文后面
	newText := append(plainText, addPlain...)
	return newText
}

func main() {
	//des的cbc加密

	a := 10
	b := a ^ 0
	fmt.Println(b)

	fmt.Println("cbc")
}
