package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

//定义一个工作量证明的结构ProofWork
type ProofWork struct {
	block  *Block
	target *big.Int
}

//2.提供创建POW的函数
func NewProofWork(block *Block) *ProofWork {
	pw := ProofWork{
		block: block,
	}
	//我们指定的难度值，现在是一个string类型，需要进行转换
	targetStr := "0000100000000000000000000000000000000000000000000000000000000000"
	//引入的辅助变量，目的是将上面的难度值转成big.int
	tmpInt := big.Int{}
	//将难度值赋值给big.int，指定16进制的格式
	tmpInt.SetString(targetStr, 16)
	pw.target = &tmpInt
	return &pw
}

//3. 提供计算 不断计算hash的函数
func (pw *ProofWork) Run() ([]byte, uint64) {
	//1.拼装数据（区块数据 还有不断变换对随机数）
	//2. 做hash运算
	//3.与pw中对target进行比较
	//找到了 退出返回
	//没找到 继续找，随机数+1

	var nonce uint64
	block := pw.block
	var hash [32]byte
	for {
		//1.拼装数据
		bytesArray := [][]byte{
			Uint64ToByte(block.Version),
			block.PrevHash,
			block.MerkelRoot,
			Uint64ToByte(block.TimeStamp),
			Uint64ToByte(block.Difficulty),
			Uint64ToByte(nonce),
			block.Data,
		}
		//将一个二维的切片连接起来，生成一唯的切片
		resultData := bytes.Join(bytesArray, []byte(""))

		//2 hash运算
		hash = sha256.Sum256(resultData)

		//3.比较
		tempInt := big.Int{}
		//将hash转成一个big.int
		tempInt.SetBytes(hash[:])
		//比较当前hash与目标对hash，如果当前hash小于目标hash，就是找到了

		if tempInt.Cmp(pw.target) == -1 {
			fmt.Printf("挖矿成功! %x; nonce: %d\n", hash, nonce)
			break
		} else {
			nonce++
		}
	}

	//最后得到两个数，一个hash，一个随机值
	return hash[:], nonce
}
