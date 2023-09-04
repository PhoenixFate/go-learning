package main

import "fmt"

func testA2() {
	fmt.Println("aaa")
}

/*
如果调用了内置函数recover，
并且定义该defer语句的函数发生了panic异常，
recover会使程序从panic中恢复，并返回panic value。
导致panic异常的函数不会继续运行，但能正常返回。
在未发生panic时调用recover，recover会返回nil。
*/
func testB2(index int) {
	//设置recover
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println("bbb")
	//显示调用panic函数，会中断程序
	//panic("中断程序 after testB")
	//隐士调用
	var a [10]int
	fmt.Println(a[index])
}

func testC21() {
	//延迟调用中引发的错误，可被后续延迟调用捕获，但仅最后⼀个错误可被捕获：
	defer func() {
		fmt.Println(recover())
	}()
	defer func() {
		panic("defer panic")
	}()
	panic("test panic")
}

func testC2() {
	fmt.Println("ccc")
}

func main() {
	//使用panic 会导致程序中断
	//使用recover继续程序
	testA2()
	testB2(29)
	testC21()
	testC2()

	fmt.Println("03 recover")
}
