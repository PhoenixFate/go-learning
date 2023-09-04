package main

import "fmt"

func main() {

	var s1 []int
	fmt.Printf("s1 cap: %d; s1 len:%d\n", cap(s1), len(s1))
	//以两倍的容量来扩容
	s1 = append(s1, 1)
	fmt.Printf("s1 cap: %d; s1 len:%d\n", cap(s1), len(s1))
	s1 = append(s1, 2)
	fmt.Printf("s1 cap: %d; s1 len:%d\n", cap(s1), len(s1))
	s1 = append(s1, 3)
	fmt.Printf("s1 cap: %d; s1 len:%d\n", cap(s1), len(s1))
	s1 = append(s1, 3)
	fmt.Printf("s1 cap: %d; s1 len:%d\n", cap(s1), len(s1))
	s1 = append(s1, 5)
	fmt.Printf("s1 cap: %d; s1 len:%d\n", cap(s1), len(s1))

	fmt.Println("11 slice append")
}
