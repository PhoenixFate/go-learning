package main

import "fmt"

func mapDelete(m map[int]string, key int) {
	delete(m, key) //删除m中 键值为key的map元素 如果删除不存在的key 不会报错
}

func main() {
	m1 := map[int]string{1: "tom", 2: "acb", 3: "lucy"}
	fmt.Println(m1)
	mapDelete(m1, 3)
	fmt.Println(m1)
}
