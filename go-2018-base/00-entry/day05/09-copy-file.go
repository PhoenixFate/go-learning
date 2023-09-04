package main

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(src string, dest string) {
	//打开源文件
	srcFile, error1 := os.Open(src)
	if error1 != nil {
		fmt.Println("open src error: ", error1)
		return
	}
	//新建目标文件
	destFile, error2 := os.Create(dest)
	if error2 != nil {
		fmt.Println("create dest error: ", error2)
		return
	}
	//关闭文件
	defer srcFile.Close()
	defer destFile.Close()

	//拷贝文件
	buffer := make([]byte, 1024)
	for {
		n, error3 := srcFile.Read(buffer)
		if error3 != nil {
			if error3 == io.EOF {
				break
			}
			fmt.Println("read file error: ", error3)
			break
		}
		//往目标文件写内容，读多少写多少
		destFile.Write(buffer[:n])
	}
}

func main() {
	src := "./1.jpg"
	dest := "./2.jpg"
	CopyFile(src, dest)

	fmt.Println("09 拷贝文件")
}
