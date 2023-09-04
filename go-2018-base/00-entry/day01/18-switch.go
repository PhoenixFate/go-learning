package main

import "fmt"

func main() {
	num := 1
	switch num {
	case 1:
		fmt.Println("1")
		//go 默认就是break
		//但可以使用fallthrough 强制调用后面的代码
		fallthrough
	case 2:
		fmt.Println("2")
		//go 默认就是break
		break
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("default")
	}

	//支持一个初始化语句，初始化语句和变量本身，以分号分隔
	switch num2 := 1; num2 {
	case 1:
		fmt.Println("num2 1")
		//go 默认就是break
		//但可以使用fallthrough 强制调用后面的代码
		fallthrough
	case 2:
		fmt.Println("num2 2")
		//go 默认就是break
		break
	case 3:
		fmt.Println("num23")
	default:
		fmt.Println("default")
	}

	var score int
	score = 85
	//switch 可以没有条件
	switch {
	case score >= 85:
		fmt.Println("优秀")
	case score >= 75:
		fmt.Println("良好")
	case score >= 60:
		fmt.Println("合格")
	default:
		fmt.Println("不及格")
	}

	fmt.Println("18 switch")
}
