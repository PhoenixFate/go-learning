package main

import (
	"fmt"
	"runtime"
)

func main() {

	//返回的n是上次设置的核数
	n := runtime.GOMAXPROCS(1)
	fmt.Println("n=", n)

	for {
		go fmt.Print(1)
		fmt.Print(0)
	}

	fmt.Println("04 设置并行计算的cpu最大核数")
}
