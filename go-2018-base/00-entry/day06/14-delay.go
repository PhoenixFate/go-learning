package main

import (
	"fmt"
	"time"
)

func delay1() {
	time.Sleep(time.Second * 2)
	fmt.Println("time.sleep")
}

func delay2() {
	timer := time.NewTimer(time.Second * 2)
	<-timer.C
	fmt.Println("time.NewTimer")
}

func delay3() {
	//time.After返回一个channel
	ch := time.After(time.Second * 2)
	t := <-ch
	fmt.Println("t: ", t)
	fmt.Println("time.After")
}

func main() {
	delay1()
	delay2()
	delay3()
	fmt.Println("14 延时的几种方式")
}
