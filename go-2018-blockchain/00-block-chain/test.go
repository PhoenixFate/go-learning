package main

import (
	"crypto/sha256"
	"fmt"
	"reflect"
)

type Test struct {
	name string
	age  int
}

func testParameter(t Test) {
	fmt.Printf("test address:%p\n", &t)
	fmt.Printf("name address: %p\n", &t.name)
}

func testStringSlice(s []string) {
	fmt.Printf("test string slice: %T\n", s)
}

func main() {
	b := make([]byte, 10)
	b[0] = 'a'
	b[1] = 'c'
	b[2] = 'm'
	a := 10
	c := &a
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(b[:]))
	fmt.Println(string(b))
	fmt.Println(string(b[:]))
	fmt.Println(string(append(b, 'i', 'l', 'v', 'u')))
	fmt.Println(string(b[:]))

	hash := sha256.Sum256(b[:])
	//ed5cfbf9a26aad64ffcf76952c1a0232cfa206b1c779795c0788183815ad175e
	//对于任意长度的消息，SHA256都会产生一个256bit长的哈希值，称作消息摘要。
	//这个摘要相当于是个长度为32个字节的数组，通常用一个长度为64的十六进制字符串来表示
	fmt.Printf("hash: %x\n", hash)

	t := Test{name: "abc", age: 20}
	fmt.Printf("t address: %p\n", &t)
	fmt.Printf("name address: %p\n", &t.name)
	testParameter(t)

	s := "string"
	fmt.Printf("s type:%s %T\n", s, s)
	fmt.Printf("s[:] type:%s %T\n", s, s[1:len(s)-1])
	ss := s[:]
	fmt.Printf("ss type: %s %T", ss, ss)
	//testStringSlice(ss[:])
	//log.Panic("log panic test")
}
