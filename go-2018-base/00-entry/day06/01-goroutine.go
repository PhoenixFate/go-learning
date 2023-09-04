package main

import "fmt"
import "time"

func NewRoutine() {
	for i := 0; i < 20; i++ {
		fmt.Println("new routine")
		time.Sleep(time.Second)
	}
}

//主协程退出，其他子协程也会退出
func main() {
	//使用go 关键字来调用协程
	go NewRoutine()

	for i := 0; i < 5; i++ {
		fmt.Println("04-01-firstBlockChain routine")
		time.Sleep(time.Second)
	}

	fmt.Println("01 协程 goroutine")
}
