package main

import (
	"fmt"
	"strings"
)

func wordsCount(words []string) map[string]int {
	m := make(map[string]int, 10)
	for _, word := range words {
		fmt.Println(word)
		if _, has := m[word]; has {
			m[word] = m[word] + 1
		} else {
			m[word] = 1
		}

	}
	return m
}

func main() {
	str := "I lave my work and I love my family too"
	//strings.Fields 就是根据空格 拆分字符串，返回字符串切片
	words := strings.Fields(str)
	fmt.Println(words)
	m := wordsCount(words)
	fmt.Println(m)

	for k, v := range m {
		fmt.Printf("word: %q; count:%d\n", k, v)
	}
}
