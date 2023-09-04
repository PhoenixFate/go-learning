package main

import "calculate"
import "fmt"

//只需要导入这个包，但不需要引用这个但包，可以使用空白表示符_ 来重命名这个导入
import _ "test"

func init() {
	fmt.Println("main package init")
}

//需要在GoLand中设置GoPath
func main() {
	result := calculate.Add(10, 20)
	fmt.Println("result= ", result)

	fmt.Println("不同文件的包管理")
}
