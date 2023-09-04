package main

import "fmt"

func swap(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	a := 10
	b := 20
	swap(&a, &b)
	fmt.Printf("a=%d, b=%d\n", a, b)
}
