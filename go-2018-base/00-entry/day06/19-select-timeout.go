package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case num := <-ch:
				fmt.Println("num: ", num)
			case <-time.After(time.Second * 3):
				fmt.Println("超时")
				quit <- true
				return
			}
		}
	}()

	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Second)
	}

	time.Sleep(time.Second * 10)

	fmt.Println("19 select超时检测")
}
