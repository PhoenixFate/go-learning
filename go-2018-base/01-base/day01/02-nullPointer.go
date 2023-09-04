package main

import "fmt"

func main() {

	//空指针
	var p *int
	// runtime error: invalid memory address or nil pointer dereference
	*p = 100
	fmt.Println(*p)

	//野指针 指针指向一块无效的内存区
	//var pp *int = 0
	//*pp = 1000
}
