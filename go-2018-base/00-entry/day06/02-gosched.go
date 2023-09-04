package main

import (
	"fmt"
	"runtime"
)

func main() {

	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Println("new routine ", i)
		}
	}()

	for i := 0; i < 5; i++ {
		//让出时间片
		runtime.Gosched()
		fmt.Println("04-01-firstBlockChain routine ", i)
	}
	//runtime.Gosched()
	fmt.Println("02 gosched 让出时间片")
}
