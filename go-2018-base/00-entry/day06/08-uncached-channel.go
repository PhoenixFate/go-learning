package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 0)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
			fmt.Printf("子协程[%d]:  len:%d, cap:%d\n", i, len(ch), cap(ch))
		}
	}()

	//延时
	time.Sleep(time.Second * 2)

	for i := 0; i < 3; i++ {
		num := <-ch
		fmt.Println("num: ", num)
	}

	fmt.Println("08 无缓存的channel")
}
