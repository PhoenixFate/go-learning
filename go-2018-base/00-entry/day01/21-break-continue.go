package main

import "fmt"

func main() {

	for i := 1; ; i++ {
		fmt.Println(i)
		if i > 20 {
			//break 可用于for、switch、select
			break
		} else {
			//而continue仅能用于for循环
			continue
		}
	}

	fmt.Println("21 break continue")
}
