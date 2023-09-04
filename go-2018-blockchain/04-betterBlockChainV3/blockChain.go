package main

import (
	"github.com/boltdb/bolt"
	"log"
)

//4.引入区块链
//使用数据代替数组
//每次添加区块的时候 都要做两件事情
//1.添加区块数据
//2.更新lastHashKey的value
type BlockChain struct {
	//区块链数组
	//blocks []*Block
	db   *bolt.DB
	tail []byte //存储最后一个区块的hash
}

const blockChainDb = "blockChain.db"
const blockBucket = "blockBucket"

//5.定义一个区块链
func NewBlockChain() *BlockChain {
	//1.打开数据库  0600 读写
	db, openDbError := bolt.Open(blockChainDb, 0600, nil)
	//defer db.Close()
	if openDbError != nil {
		log.Panic("开始数据库失败")
	}
	var lastHash []byte //最后一个区块的hash

	//2.找到抽屉bucket，如果没有则创建
	_ = db.Update(func(tx *bolt.Tx) error {
		//tx 事务
		// Returns nil if the bucket does not exist.
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			//不存在则创建bucket
			var createBucketError error
			bucket, createBucketError = tx.CreateBucket([]byte(blockBucket))
			if createBucketError != nil {
				log.Panic("创建bucket test失败")
			}
			//创建一个创世块，并作为第一个区块添加到区块链中
			genesisBlock := GenesisBlock()
			//hash作为key，block的字节流作为value
			_ = bucket.Put(genesisBlock.Hash, genesisBlock.Serialize())
			_ = bucket.Put([]byte("lastHashKey"), genesisBlock.Hash)
			lastHash = genesisBlock.Hash
		} else {
			lastHash = bucket.Get([]byte("lastHashKey"))
		}
		//已经有bucket了
		return nil
	})
	return &BlockChain{db, lastHash}
}

func GenesisBlock() *Block {
	return NewBlock("创世快", []byte{})
}

//6.添加区块
func (blockChain *BlockChain) AddBlock(data string) {
	//获得最后一个区块的hash
	db := blockChain.db
	lastHash := blockChain.tail
	_ = db.Update(func(tx *bolt.Tx) error {
		//tx 事务
		// Returns nil if the bucket does not exist.
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			log.Panic("bucket 为空")
		}
		//6.1创建新的区块
		block := NewBlock(data, lastHash)
		//6.2添加到区块链数据中
		_ = bucket.Put(block.Hash, block.Serialize())
		_ = bucket.Put([]byte("lastHashKey"), block.Hash)
		blockChain.tail = block.Hash
		return nil
	})
}
