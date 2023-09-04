package calculate

import "fmt"

func init() {
	fmt.Println("this is calculate package init")
}

func Add(a, b int) int {
	return a + b
}
