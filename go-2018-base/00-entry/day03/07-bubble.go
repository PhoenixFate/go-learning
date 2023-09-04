package main

import "fmt"
import "math/rand"
import "time"

func main() {
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UnixNano())
	rand.Seed(time.Now().Unix())
	a := [10]int{}
	for i := 0; i < len(a); i++ {
		a[i] = rand.Intn(100)
	}

	//冒泡排序
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	fmt.Println(a)

	fmt.Println("07 冒泡")
}
