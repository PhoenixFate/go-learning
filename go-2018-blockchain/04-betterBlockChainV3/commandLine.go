package main

import (
	"fmt"
	"os"
)

//这是一个用来接受命令行参数并且控制区块链操作的文件
type CommandLine struct {
	bc *BlockChain
}

const Usage = `
	addBlock --data DATA "add data to blockchain"
	printChain "print all blockchain data"
`

//接受参数的动作，我们放到一个函数中
func (commandLine *CommandLine) Run() {
	//1.得到所有的命令
	args := os.Args
	fmt.Println(args)
	if len(args) < 2 {
		fmt.Println("参数过少；请参考下面使用说明")
		fmt.Printf(Usage)
		return
	}
	//2.分析命令
	cmd := args[1]
	//3.根据不同的命令，执行相应的动作
	switch cmd {
	case "addBlock":
		//添加区块
		fmt.Println("add block")
		//确保命令有效
		if len(args) == 4 && args[2] == "--data" {
			//获取命令行的数据
			data := args[3]
			commandLine.AddBlock(data)
		} else {
			fmt.Println("添加区块 参数使用不当，请检查")
		}
	case "printChain":
		fmt.Println("print chain")
		commandLine.PrintBlockChain()
	default:
		fmt.Printf(Usage)
	}
}

func (commandLine *CommandLine) AddBlock(data string) {
	commandLine.bc.AddBlock(data)
}

func (commandLine *CommandLine) PrintBlockChain() {
	it := commandLine.bc.NewIterator()
	//调用迭代器
	for {
		//Next 返回当前区块，并且指针左移
		block := it.Next()
		fmt.Printf("===========================\n")
		fmt.Printf("前区块哈希值：%x\n", block.PrevHash)
		fmt.Printf("区块哈希值： %x\n", block.Hash)
		fmt.Printf("区块数据： %s\n", block.Data)
		fmt.Printf("时间戳： %d\n", block.TimeStamp)
		fmt.Printf("随机数: %d\n", block.Nonce)
		fmt.Printf("难度值: %d\n", block.Difficulty)
		fmt.Printf("===========================\n\n")
		if len(block.PrevHash) == 0 {
			fmt.Println("区块链编译结束")
			break
		}
	}
}
