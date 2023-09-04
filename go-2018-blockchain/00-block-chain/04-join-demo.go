package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	strArr := []string{"hello", "world", "!"}
	result := strings.Join(strArr, "")
	fmt.Println(result)

	secondByteArray := make([][]byte, 3)
	secondByteArray[0] = []byte("hello")
	secondByteArray[1] = []byte("world")
	secondByteArray[2] = []byte("!")
	result2 := bytes.Join(secondByteArray, []byte(""))
	fmt.Println(string(result2))
}
