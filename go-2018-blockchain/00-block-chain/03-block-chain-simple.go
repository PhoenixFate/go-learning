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

//4.引入区块链
type BlockChain struct {
	//区块链数组
	blocks []*Block
}

//5.定义一个区块链
func NewBlockChain() *BlockChain {
	//创建一个创世块，并作为第一个区块添加到区块链中
	genesisBlock := GenesisBlock()
	return &BlockChain{
		blocks: []*Block{genesisBlock},
	}
}

func GenesisBlock() *Block {
	return NewBlock("创世快", []byte{})
}

//6.添加区块
func (blockChain *BlockChain) AddBlock(data string) {
	//获得最后一个区块
	lastBlock := blockChain.blocks[len(blockChain.blocks)-1]
	prevHash := lastBlock.Hash

	//6.1创建新的区块
	block := NewBlock(data, prevHash)
	//6.2添加到区块链数据中
	blockChain.blocks = append(blockChain.blocks, block)
}

//6.重构代码

func main() {
	blockChain := NewBlockChain()
	blockChain.AddBlock("班长向班花转了50枚比特币")
	for i, block := range blockChain.blocks {
		fmt.Printf("当前区块高度：%d\n", i)
		fmt.Printf("前区块哈希值：%x\n", block.PrevHash)
		fmt.Printf("当前区块哈希值： %x\n", block.Hash)
		fmt.Printf("区块数据： %s\n", block.Data)
	}
	//block := NewBlock("送给你一枚比特币", []byte{})

}
