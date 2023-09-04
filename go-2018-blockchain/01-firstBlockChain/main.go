package main

import (
	"fmt"
)

//cd 00-block-chain/01-firstBlockChain
//go run *.go

func main() {
	blockChain := NewBlockChain()
	blockChain.AddBlock("班长向班花转了50枚比特币")
	for i, block := range blockChain.blocks {
		fmt.Printf("当前区块高度：%d\n", i)
		fmt.Printf("前区块哈希值：%x\n", block.PrevHash)
		fmt.Printf("当前区块哈希值： %x\n", block.Hash)
		fmt.Printf("区块数据： %s\n", block.Data)
	}
	var s string
	fmt.Scanf("%s", s)
}
