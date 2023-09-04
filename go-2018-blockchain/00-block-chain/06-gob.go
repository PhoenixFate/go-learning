package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

//gob是golang 内置的包 用于序列号 反序列化
type Person struct {
	Name string
	Age  uint
}

func main() {

	var p Person
	p.Name = "tom"
	p.Age = 20

	//编码的数据放到buffer中
	var buffer bytes.Buffer

	//使用gob进行序列化（编码）得到字节流
	//1.定义一个编码器
	//2.使用编码器进行编码
	encoder := gob.NewEncoder(&buffer)
	error1 := encoder.Encode(&p)
	if error1 != nil {
		log.Panic("编码person失败", error1)
	}
	fmt.Printf("编码后的person%v\n", buffer.Bytes())

	//使用gob进行序列化（解码）得到Person结构
	decoder := gob.NewDecoder(bytes.NewReader(buffer.Bytes()))
	var newPerson Person
	decodeError := decoder.Decode(&newPerson)
	if decodeError != nil {
		log.Panic("解码person失败", decodeError)
	}
	fmt.Println("解码后的person: ", newPerson)
	//1.定义一个解码器
	//2.使用编码器进行编码

}
