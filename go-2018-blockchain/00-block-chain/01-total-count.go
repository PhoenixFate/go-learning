package main

import "fmt"

func main() {
	//1btc=10^8sat  (Satoshi 中本聪)
	//1. 每21万个块减半
	//2. 最初奖励50个btc  bitcoin
	//3. 循环累加
	fmt.Println("total count")
	total := 0.0
	blockInterval := 21 //单位万
	currentAward := 50.0
	reduceCount := 0 //衰减次数
	for currentAward > 0.00000001 {
		amount1 := float64(blockInterval) * currentAward
		currentAward *= 0.5 //使用乘法来代替除法
		total += amount1
		reduceCount++
	}
	fmt.Println("预计到 ", 2009+4*reduceCount, "年所有矿挖完")
	fmt.Println("比特币总量", total, "万")
}
