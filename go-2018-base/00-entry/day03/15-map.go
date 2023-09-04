package main

import "fmt"

//map传递的是引用
func myTest(m map[int]string) {
	delete(m, 2)
}

func main() {
	//map[keyType]valueType
	var map1 map[int]string //只是声明没有初始化
	fmt.Println(map1)
	fmt.Println(map1 == nil) //true
	fmt.Println("length map1: ", len(map1))
	//map1[0]="hello" //err, panic: assignment to entry in nil map

	//通过make创建map
	map10 := map[int]string{} //声明并且初始化
	map10[0] = "hello"
	fmt.Println("map10: ", map10)
	var map2 = make(map[string]string) //这个和上面那个方法是等价的
	map2["a"] = "aaa"
	map2["b"] = "bbb"
	map2["c"] = "ccc"
	fmt.Println(map2)
	fmt.Println("length map2: ", len(map2))

	//make创建map，指定容量
	//但不够用，也会自动扩容
	var map3 = make(map[string]string, 10)
	fmt.Println(map3)
	fmt.Println("length map3: ", len(map3))
	//fmt.Println("capt map3: ", cap(map3)) //cap不能传入map

	//定义的时候初始化
	var map4 = map[int]string{1: "abc", 2: "bcd", 3: "edf"}
	fmt.Println(map4)
	//修改某个key对应的值
	map4[1] = "ggg"
	fmt.Println(map4)

	//遍历map
	for key, value := range map4 {
		fmt.Printf("key: %d; value:%s\n", key, value)
	}

	//判断map中的key是否存在
	value, ok := map4[0]
	if ok {
		fmt.Println("key 存在, value: ", value)
	} else {
		fmt.Println("key 不存在")
	}

	//map删除某一个key
	delete(map4, 1)
	fmt.Println(map4)

	myTest(map4)
	fmt.Println(map4)

	//map作为参数传递是引用传递
	fmt.Println("--- 15 map ---")
}
