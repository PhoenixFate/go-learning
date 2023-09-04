package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(time.Second * 2)
	go func() {
		<-timer.C
		fmt.Println("定时器时间到了")
	}()

	//停止计时器
	timer.Stop()

	timer2 := time.NewTimer(time.Second * 1)
	timer2.Reset(time.Second * 3)

	<-timer2.C
	fmt.Println("timer2 reset")

	time.Sleep(time.Second * 10)

	fmt.Println("15 停止 重置定时器")
}
