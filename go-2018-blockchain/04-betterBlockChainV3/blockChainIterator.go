package main

import (
	"github.com/boltdb/bolt"
	"log"
)

type BlockChainIterator struct {
	db *bolt.DB
	//游标  用于不断索引
	currentHashPointer []byte
}

func (blockChain *BlockChain) NewIterator() *BlockChainIterator {
	return &BlockChainIterator{
		db: blockChain.db,
		//最初指向区块链的最后一个区块，随着Next的调用，不断变化
		currentHashPointer: blockChain.tail,
	}
}

//迭代器是属于区块链的
//Next方法是属于迭代器的
//返回1. 当前的区块
//2.指针前移
func (it *BlockChainIterator) Next() *Block {
	var block *Block
	it.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			log.Panic("迭代器遍历时 bucket 为空")
		}
		blockTemp := bucket.Get(it.currentHashPointer)
		//解码
		block = Deserialize(blockTemp)
		//指针左移
		it.currentHashPointer = block.PrevHash
		return nil
	})

	return block
}
