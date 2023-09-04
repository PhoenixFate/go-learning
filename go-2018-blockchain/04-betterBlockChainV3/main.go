package main

//mac/linux 下面
//cd 00-block-chain/01-firstBlockChain
//go run *.go

//windows下面编译运行
//go build
//start xxx.exe
//记得代码后面+ time.Sleep(time.Second*100) or fmt.Scanf()

func main() {
	blockChain := NewBlockChain()
	commandLine := CommandLine{blockChain}
	commandLine.Run()
	//time.Sleep(time.Second * 100)

}
