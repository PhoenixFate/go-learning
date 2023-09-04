package main

import "os"
import "fmt"

//在终端调用该文件
//go run 09-04-01-firstBlockChain-args.go a mike
func main() {

	//接受用户传递过来的参数，都是以字符串传递过来的
	list := os.Args
	for i := 0; i < len(list); i++ {
		fmt.Println(list[i])
	}
	fmt.Println("09 接受命令行传递过来的参数")
}
