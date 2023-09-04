package main

import "fmt"

func main() {
	var a int
	a = 10
	var b = 20
	if a > b {
		fmt.Println("a > b")
	} else {
		fmt.Println("a < b")
	}

	//if支持1个初始化语句，初始化语句和判断条件以分号分隔
	if c := 10; c == 10 {
		fmt.Println("c==10")
	} else {
		fmt.Println("c!=10")
	}

	if a > b {
		fmt.Println("a > b")
	} else if a == b {
		fmt.Println("a == b")
	} else {
		fmt.Println("a < b")
	}

	fmt.Println("17 if")
}
