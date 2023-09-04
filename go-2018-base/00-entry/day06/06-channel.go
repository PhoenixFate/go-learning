package main

import (
	"fmt"
	"time"
)

//定义一个全局channe
var ch = make(chan string)

//定义打印机，模拟多个协程同时使用打印机
func PrintMachine6(content string) {

	for _, data := range content {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Println()
}

func Person61() {
	PrintMachine6("hello")
	//给管道发送数据
	ch <- "sss"
}

func Person62() {
	//从管道取数据，如果通道没有数据，就会阻塞等待
	<-ch
	PrintMachine6("world")
}

func main() {

	go Person61()
	go Person62()

	time.Sleep(time.Second * 10)
	fmt.Print("06 channel")
}
