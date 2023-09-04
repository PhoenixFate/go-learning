package main

import "fmt"
import "os"

func main() {
	//可以手动关闭默认的标准输出设备
	//os.Stdout.Close()
	//标准输出设备
	//fmt.Println 往标准输出设备打印消息，标准输出设备默认是打开的
	fmt.Println("are you ok?")
	_, _ = os.Stdout.WriteString("os stdout write string\n")

	//关闭标准输入文件
	//os.Stdin.Close()
	//标准输入设备
	var a int
	fmt.Println("请输入 a 的值")
	fmt.Scan(&a)
	fmt.Println("刚刚输入的a 的值：", a)

	fmt.Println("07 stdin stdout")
}
