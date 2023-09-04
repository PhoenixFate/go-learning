package main

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"log"
	"time"
)

//1. 定义区块结构
type Block struct {
	//1.版本号
	Version uint64
	//2. 前区块hash
	PrevHash []byte
	//3.Merkel根
	MerkelRoot []byte
	//4.时间戳
	TimeStamp uint64
	//5.难度值
	Difficulty uint64
	Nonce      uint64
	//a. 当前区块hash 正常比特币区块中没有当前区块的哈希
	Hash []byte
	//b. 数据
	Data []byte
}

func (block *Block) Serialize() []byte {
	var buffer bytes.Buffer
	//使用gob进行序列化（编码）得到字节流
	//1.定义一个编码器
	//2.使用编码器进行编码
	encoder := gob.NewEncoder(&buffer)
	encodeError := encoder.Encode(&block)
	if encodeError != nil {
		log.Panic("编码block失败", encodeError)
	}
	return buffer.Bytes()
}

func Deserialize(data []byte) *Block {
	//使用gob进行序列化（解码）得到Person结构
	//1.定义一个解码器
	//2.使用编码器进行编码
	decoder := gob.NewDecoder(bytes.NewReader(data))
	var block Block
	decodeError := decoder.Decode(&block)
	if decodeError != nil {
		log.Panic("解码block失败", decodeError)
	}
	return &block
}

//2.创建区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		Version:    00,
		PrevHash:   prevBlockHash,
		MerkelRoot: []byte{},
		TimeStamp:  uint64(time.Now().Unix()),
		Difficulty: 0,
		Nonce:      0,
		Hash:       []byte{}, //先空，后面计算hash
		Data:       []byte(data),
	}
	//block.SetHash()
	//创建一个ProofWork对象
	pw := NewProofWork(&block)
	//查找随机数 不停的进行hash运算
	hash, nonce := pw.Run()
	//根据挖矿结果对区块结果 进行更新（补充）
	block.Hash = hash
	block.Nonce = nonce
	return &block
}

//辅助函数 将uint64 转[]byte
func Uint64ToByte(num uint64) []byte {
	var buffer bytes.Buffer
	//binary.BigEndian 大端对齐：高位在前 底位在后
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buffer.Bytes()
}

//3.生成hash
//已经使用工作量 证明替换setHash
/**
func (block *Block) SetHash() {
	//1. 拼装数据 使用bytes.join 来优化下面代码
	//var blockInfo []byte
	//blockInfo = append(blockInfo, Uint64ToByte(block.Version)...)
	//blockInfo = append(blockInfo, block.PrevHash...)
	//blockInfo = append(blockInfo, block.MerkelRoot...)
	//blockInfo = append(blockInfo, Uint64ToByte(block.TimeStamp)...)
	//blockInfo = append(blockInfo, Uint64ToByte(block.Difficulty)...)
	//blockInfo = append(blockInfo, block.Hash...)
	//blockInfo = append(blockInfo, block.Data...)

	bytesArray := [][]byte{
		Uint64ToByte(block.Version),
		block.PrevHash,
		block.MerkelRoot,
		Uint64ToByte(block.TimeStamp),
		Uint64ToByte(block.Difficulty),
		block.Data,
	}
	//将一个二维的切片连接起来，生成一唯的切片
	resultData := bytes.Join(bytesArray, []byte(""))

	//2.sha256
	hash := sha256.Sum256(resultData)
	block.Hash = hash[:]
}
*/
