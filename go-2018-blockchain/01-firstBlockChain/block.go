package main

import (
	"crypto/sha256"
	"fmt"
	"reflect"
)

//1. 定义区块结构
type Block struct {
	//1. 前区块hash
	PrevHash []byte
	//2. 当前区块hash
	Hash []byte
	//3. 数据
	Data []byte
}

//2.创建区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		PrevHash: prevBlockHash,
		Hash:     []byte{}, //先空，后面计算hash  //TODO
		Data:     []byte(data),
	}
	fmt.Println("block type: ", reflect.TypeOf(block))
	block.SetHash()

	return &block
}

//3.生成hash
func (block *Block) SetHash() {
	//1. 拼装数据
	blockInfo := append(block.PrevHash, block.Data...)
	//2.sha256
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}
