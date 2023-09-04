package main

import (
	"fmt"
	"time"
)

//定义打印机，模拟多个协程同时使用打印机
func PrintMachine(content string) {

	for _, data := range content {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Println()
}

func Person1() {
	PrintMachine("hello")
}

func Person2() {
	PrintMachine("world")

}

func main() {

	go Person1()
	go Person2()

	time.Sleep(time.Second * 10)
	fmt.Print("05 多任务 竞争同1个资源")
}
