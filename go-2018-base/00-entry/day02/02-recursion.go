package main

import "fmt"

func MyRecursion(a int) {
	if a == 1 {
		fmt.Println("a=", a)
		return
	}
	fmt.Println("a=", a)
	MyRecursion(a - 1)
}

func sumRecursion(a int) (sum int) {
	if a == 1 {
		return 1
	}
	return a + sumRecursion(a-1)
}

func main() {
	MyRecursion(10)
	sum := sumRecursion(100)
	fmt.Println("sum: ", sum)
	fmt.Println("02 recursion")
}
