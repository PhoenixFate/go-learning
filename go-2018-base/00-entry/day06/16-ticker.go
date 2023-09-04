package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.Tick(time.Second)

	i := 0
	for {
		i++
		t := <-ticker
		fmt.Println(t)
		if i == 10 {
			break
		}
	}

	fmt.Println("16 ticker 周期任务")
}
