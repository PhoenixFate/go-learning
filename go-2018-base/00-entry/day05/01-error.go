package main

import "fmt"
import "errors"

func MyDivide(a, b int) (result int, resultError error) {
	if b == 0 {
		//errors.New() 返回一个error对象
		resultError = errors.New("除数不能为0")
	} else {
		result = a / b
	}
	return
}

func main() {
	fmt.Println("-----------------------------")
	//fmt.Errorf 返回一个error对象
	error1 := fmt.Errorf("%s", "this is error 1")
	fmt.Println("error1: ", error1)

	//errors.New() 返回一个error对象
	error2 := errors.New("this is error2")
	fmt.Println("error2: ", error2)

	result, err := MyDivide(10, 2)
	fmt.Println("result: ", result)
	fmt.Println("result error: ", err)

	result2, err2 := MyDivide(10, 0)
	fmt.Println("result2: ", result2)
	fmt.Println("result error2: ", err2)

	fmt.Println("01 error")
}
