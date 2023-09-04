package main

import "fmt"

func main() {
	//创建一个双向的channel
	ch := make(chan int)
	//双向channel 能够隐式转换为单向的channel
	var writeChannel chan<- int = ch
	var readChannel <-chan int = ch

	go func() {
		writeChannel <- 666
	}()

	num := <-readChannel
	fmt.Println("num: ", num)

	fmt.Println("11 单向的 channel")
}
