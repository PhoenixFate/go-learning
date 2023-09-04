package main

import (
	"fmt"
	"os"
)

func main() {

	//1.创建文件 create
	//文件不存在则创建 文件存在则将文件内容清空
	//参数name 文件路径：相对路径、绝对路径
	fPointer, err1 := os.Create("./abc.txt")
	if err1 != nil {
		fmt.Println("create abc error: ", err1)
		return
	}
	//defer 在最后执行
	defer fPointer.Close()
	fmt.Println("create abc.txt successfully")

	//2.打开文件Open
	//以只读方式打开文件 如果文件不存在 打开失败
	fPointer2, err2 := os.Open("./bcd.txt")
	if err2 != nil {
		fmt.Println("open bcd error: ", err2)
	}
	b := make([]byte, 1024)
	fPointer2.Read(b)
	fmt.Println("read file: ", string(b))

	//3.打开文件Openfile
	//以只读、只写、读写方式打开文件
	//name 文件路径 相对路径 绝对路径
	//flag 打开文件的模式 O_RDONLY(只读模式), O_WRONLY(只写) O_RDWR(读写) O_APPEND(追加)
	//perm 权限取值范围0-7
	//     0 没有任何权限
	//     1 执行权限
	//     2 写权限
	//     3 写与执行
	//     4 读
	//     5 读与执行
	//     6 读写
	//     7 读写与执行
	fPointer3, err3 := os.OpenFile("hello.txt", os.O_RDWR, 6)
	if err3 != nil {
		fmt.Println("open file failed")
	}
	c := make([]byte, 10)
	fPointer3.Read(c)
	fmt.Println("open file read: ", string(c))
}
