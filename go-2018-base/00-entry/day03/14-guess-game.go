package main

import (
	"fmt"
	"math/rand"
	"time"
)

func createNumber(p *int) {
	//设置种子
	rand.Seed(time.Now().UnixNano())
	for {
		*p = rand.Intn(10000)
		if *p >= 1000 {
			break
		}
	}
}

func generateSlice(s []int, randomNumber int) {
	s[0] = randomNumber / 1000
	s[1] = (randomNumber % 1000) / 100
	s[2] = (randomNumber % 100) / 10
	s[3] = randomNumber % 10

}

func onGame(randomSlice []int, randomNumber int) {

	for {
		var num int
		for {
			fmt.Println("请输入一个 1000-9999之间的数组")
			_, _ = fmt.Scan(&num)
			if num >= 1000 && num < 10000 {
				break
			}
		}

		inputSlice := make([]int, 4)
		generateSlice(inputSlice, num)
		for i := 0; i < 4; i++ {
			if inputSlice[i] > randomSlice[i] {
				fmt.Printf("第%d位大了\n", i)
			} else if inputSlice[i] < randomSlice[i] {
				fmt.Printf("第%d位小了\n", i)
			} else {
				fmt.Printf("第%d位猜对了\n", i)
			}
		}
		if randomNumber == num {
			fmt.Println("恭喜你 全部猜对了，游戏结束")
			break
		}
	}

}

func main() {
	var randomInt int
	createNumber(&randomInt)
	fmt.Println(randomInt)

	numberSlice := make([]int, 4)
	//var numberSlice [4]int
	generateSlice(numberSlice, randomInt)
	fmt.Println(numberSlice)

	onGame(numberSlice, randomInt)

	fmt.Println("14 猜字游戏")
}
