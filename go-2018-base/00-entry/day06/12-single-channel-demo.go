package main

import "fmt"

//只能写 不能读
func producer(writeChannel chan<- int) {
	for i := 0; i < 3; i++ {
		writeChannel <- i * i
	}
	close(writeChannel)
}

//只能读，不能写
func consumer(readChannel <-chan int) {
	for num := range readChannel {
		fmt.Println("num: ", num)
	}
}

func main() {
	//创建一个双向的channel
	ch := make(chan int)

	//生产者，传入数据
	go producer(ch)
	//消费者，打印数据
	go consumer(ch)

	fmt.Println("12 单向channel 的demo")
}
