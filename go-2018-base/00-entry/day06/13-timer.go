package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("now: ", time.Now())
	//创建一个定时器，2s后，往time通过写内容（内容为当前时间）
	timer := time.NewTimer(2 * time.Second)

	//2s后timer会往timer通道写输入，timer.C就一个channel通道
	//timer只会写一次
	t := <-timer.C
	fmt.Println("t: ", t)

	fmt.Println("13 定时器")
}
