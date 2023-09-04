package main

import "fmt"

func testA() {
	fmt.Println("aaa")
}

func testB(index int) {
	fmt.Println("bbb")
	//显示调用panic函数，会中断程序
	//直接调用内置的panic函数也会引发panic异常
	//panic("中断程序 after testB")

	//运行时的错误，不可恢复的错误也是panic错误
	/*
		一般而言，当panic异常发生时，程序会中断运行，
		并立即执行在该goroutine（可以先理解成线程，在中被延迟的函数（defer 机制）。
		随后，程序崩溃并输出日志信息。日志信息包括panic value和函数调用的堆栈跟踪信息。
	*/
	var a [10]int
	fmt.Println(a[index])
}

func testC() {
	fmt.Println("ccc")
}

func main() {
	//panic: 严重的错误，一般当panic异常发生时，程序会中断
	//我们不应通过调用panic函数来报告普通的错误，而应该只把它作为报告致命错误的一种方式。
	testA()
	testB(20)
	testC()
	fmt.Println("02 panic")
}
