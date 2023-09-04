package main

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
