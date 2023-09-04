package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
			fmt.Printf("协程[%d] len: %d, cap:%d\n", i, len(ch), cap(ch))
		}
	}()

	time.Sleep(time.Second * 2)
	for i := 0; i < 3; i++ {
		num := <-ch
		fmt.Println("num: ", num)
	}

	fmt.Println("09 带缓存的 channel")
}
