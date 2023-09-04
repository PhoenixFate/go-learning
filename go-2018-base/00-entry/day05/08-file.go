package main

import (
	"bufio"
	"fmt"
	"io"
)
import "os"

func WriteFile(path string) {
	//打开或者新建文件
	f, error1 := os.Create(path)
	if error1 != nil {
		fmt.Println(error1)
		return
	}
	//使用完毕，关闭文件
	defer f.Close() //defer 延时调用
	for i := 0; i < 10; i++ {
		//格式化字符串
		buffer := fmt.Sprintf("i=%d\n", i)
		f.WriteString(buffer)
	}
}

func ReadFile(path string) {
	//打开文件
	f, error1 := os.Open(path)
	if error1 != nil {
		fmt.Println("open file error: ", error1)
		return
	}
	//关闭文件
	defer f.Close()

	//读文件
	buffer := make([]byte, 1024)
	//n 表示读取内容的长度
	n, error2 := f.Read(buffer)
	if error2 != nil {
		fmt.Println("read file error: ", error2)
		return
	}
	fmt.Println("buffer: ", string(buffer[:n]))
}

func ReadFileLine(path string) {
	//打开文件
	f, error1 := os.Open(path)
	if error1 != nil {
		fmt.Println("open file error: ", error1)
		return
	}
	//关闭文件
	defer f.Close()

	//新建缓冲区，把内容先放在缓冲区
	reader := bufio.NewReader(f)
	for {
		//遇到\n 就结束读取, 但是\n 也会读取进去
		buffer, error2 := reader.ReadBytes('\n')
		if error2 != nil {
			if error2 == io.EOF {
				break
			}
		}
		fmt.Printf("#%s#", buffer)
	}

}

func main() {
	path := "./demo.txt"
	WriteFile(path)
	ReadFile(path)

	ReadFileLine(path)

	fmt.Println("08 文件的操作")
}
