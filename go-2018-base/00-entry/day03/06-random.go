package main

import "fmt"
import "math/rand"
import "time"

func main() {

	//设置种子
	//如果种子一样，每次的随机数一样
	//rand.Seed(10)
	//当前系统时间作为随机数
	//js: new Date().getTime();//返回数值单位是毫秒
	//1593228608306
	//time.Now().Unix() 单位秒
	//1593228528
	fmt.Println(time.Now().Unix())
	//time.Now().UnixNano() 单位 纳秒nanoseconds
	//1593228528928911000
	fmt.Println(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())

	//产生随机数
	fmt.Println("random int: ", rand.Int())
	//某个范围内的随机数
	fmt.Println(rand.Intn(100))

	fmt.Println("06 随机数")
}
