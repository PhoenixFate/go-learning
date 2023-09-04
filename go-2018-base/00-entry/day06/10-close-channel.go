package main

import (
	"fmt"
)

func main() {

	ch := make(chan int, 0)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- i * 10
			fmt.Printf("协程[%d] len:%d, cap:%d\n", i, len(ch), cap(ch))
		}
		//如果不需要写数据，则关闭channel
		close(ch)
		//关闭channel 无法在发送数据
		//ch<-44

	}()

	//for{
	//	//如果ok为true，管道没有关闭
	//	//ok为false，管道关闭
	//	if num,ok:=<-ch;ok==true{
	//		time.Sleep(time.Second)
	//		fmt.Println("num: ",num)
	//	}else {
	//		//管道关闭
	//		break
	//	}
	//}

	//这两句等于上面注掉的
	for num := range ch {
		fmt.Println("range num: ", num)
	}

	fmt.Println("10 关闭 channel")
}
