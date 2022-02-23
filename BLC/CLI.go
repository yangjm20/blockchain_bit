package BLC

import (
	"flag"
	"fmt"
	"github.com/fabric_assert/blockchain_bit/pkg/log"
	"os"
)

//CLI 结构
type CLI struct {
	BC *BlockChain
}

//展示用法
func PrintUsage()  {
	fmt.Println("Usage:")
	fmt.Printf("\tcreateblcwithgenesis  --  创建区块链.\n")
	fmt.Printf("\taddblock -data  DATA  --  交易数据.\n")
	fmt.Printf("\tprintblockchain --  输出区块链的信息.\n")
}

//校验，如果只输入了程序命令，就输出指令用法并且退出程序
func IsValidArgs()  {
	if len(os.Args) < 2{
		PrintUsage()
		os.Exit(1)
	}
}

//添加区块
func (cli *CLI)addBlock(data string) {
	cli.BC.AddBlock([]byte(data))
}

//输出区块链信息

func (cli *CLI)printchain() {
	cli.BC.PrintChain()
}

//创建创世区块
func (cli *CLI)createBlockchinWithGenesis() {
	CreateBlockChainWithGenesisBlock()
}

//运行函数
func (cli *CLI)Run(){
	//1. 检测参数数量
	IsValidArgs()

	//2. 新建命令
	addBlockCmd:=flag.NewFlagSet("addblock",flag.ExitOnError)
	printChainCmd:=flag.NewFlagSet("printchain",flag.ExitOnError)
	createBlcWithGenesisCmd:=flag.NewFlagSet("createblockchain",flag.ExitOnError)

	//3. 获取命令行参数
	flagAddBlockArg:=addBlockCmd.String("data","send 100 BTC to everyone","交易数据")
	//flagPrintChainArg:=printChainCmd.String("data","send 100 BTC to everyone","交易数据")
	//flagcreateBlcWithGenesisArg:=createBlcWithGenesisCmd.String("data","send 100 BTC to everyone","交易数据")
	switch os.Args[1] {
	case "addblock":
		err:=addBlockCmd.Parse(os.Args[2:])
		if err!=nil{
			log.Panicf("parse cmd of addblock failed! %v\n",err.Error())
		}
	case "printchain":
		err:=printChainCmd.Parse(os.Args[2:])
		if err!=nil{
			log.Panicf("parse cmd of printchain failed! %v\n",err.Error())
		}
	case "createblockchain":
		err:=createBlcWithGenesisCmd.Parse(os.Args[2:])
		if err!=nil{
			log.Panicf("parse cmd of createblockchain failed! %v\n",err.Error())
		}
	default:
		PrintUsage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed(){
		if *flagAddBlockArg == ""{
			PrintUsage()
			os.Exit(1)
		}

		cli.addBlock(*flagAddBlockArg)
	}

	if printChainCmd.Parsed(){
		cli.printchain()
	}

	if createBlcWithGenesisCmd.Parsed(){
		cli.createBlockchinWithGenesis()
	}

}