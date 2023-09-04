package main

import (
	"fmt"
	"runtime"
	"time"
)

func MyTest5() {

	defer fmt.Println("ccccc")

	//终止所在的协程
	runtime.Goexit()

	fmt.Println("dddddd")
}

func main() {

	//创建新的协程
	go func() {
		fmt.Println("aaaaa")

		MyTest5()

		fmt.Println("bbbbb")
	}()

	//不让协程结束
	time.Sleep(time.Second * 10)

	fmt.Println("03 go exit")
}
