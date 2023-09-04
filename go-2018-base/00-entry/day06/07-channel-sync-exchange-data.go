package main

import (
	"fmt"
	"time"
)

func main() {
	//创建channel
	var ch = make(chan string)

	go func() {
		defer fmt.Println("子协程调用完毕")
		for i := 0; i < 10; i++ {
			fmt.Println("子协程 ", i)
			time.Sleep(time.Second)
		}
		ch <- "abd"
	}()

	//没有数据前会阻塞
	str := <-ch
	fmt.Println(str)
	fmt.Println("07 channel 同步以及 交换数据")
}
