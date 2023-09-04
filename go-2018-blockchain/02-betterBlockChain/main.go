package main

import (
	"fmt"
	"time"
)

//mac/linux 下面
//cd 00-block-chain/01-firstBlockChain
//go run *.go

//windows下面编译运行
//go build
//start xxx.exe
//记得代码后面+ time.Sleep(time.Second*100) or fmt.Scanf()

func main() {
	blockChain := NewBlockChain()
	blockChain.AddBlock("班长向班花转了50枚比特币")
	for i, block := range blockChain.blocks {
		fmt.Printf("当前区块高度：%d\n", i)
		fmt.Printf("前区块哈希值：%x\n", block.PrevHash)
		fmt.Printf("当前区块哈希值： %x\n", block.Hash)
		fmt.Printf("区块数据： %s\n", block.Data)
	}
	time.Sleep(time.Second * 100)
}
