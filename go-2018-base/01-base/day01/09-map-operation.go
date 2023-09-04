package main

import "fmt"

func main() {

	//遍历map
	m1 := map[int]string{1: "tom", 2: "acb", 3: "lucy"}
	for k, v := range m1 {
		fmt.Printf("key: %d; value: %s\n", k, v)
	}

	//range 返回的key value，默认可以省略value
	for k := range m1 {
		fmt.Println("key: ", k)
	}

	//使用_ 忽略key
	for _, value := range m1 {
		fmt.Println("value: ", value)
	}

	//判断map中的key是否存在
	//借助下标运算
	//m1[下标] 返回两个值，第一个是value，第二个是bool 代表key是否存在
	if v, ok := m1[12]; ok {
		fmt.Println("存在 value=", v, "ok: ", ok)
	} else {
		fmt.Println("不存在 value=", v, "ok: ", ok)
	}

}
