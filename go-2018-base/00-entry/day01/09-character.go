package main

import (
	"fmt"
)

func main() {
	var ch byte
	ch = 97

	fmt.Printf("character: %c %d\n", ch, ch)
	ch = 'b'
	fmt.Printf("character: %c %d\n", ch, ch)
	//大小写转换
	fmt.Printf("大写转小写: %c\n", 'A'+32)

	//var cc byte
	cc := '中'
	fmt.Printf("cc:%c\n", cc)
	fmt.Printf("type cc :%T\n", cc)

	dd := "中"
	fmt.Printf("dd:%s\n", dd)
	fmt.Printf("type dd :%T\n", dd)
}
