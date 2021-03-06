package BLC

import (
	"flag"
	"fmt"
	"github.com/fabric_assert/blockchain_bit/pkg/log"
	"github.com/fabric_assert/blockchain_bit/pkg/utils"
	"os"
)

//CLI 结构
type CLI struct {
	BC *BlockChain
}

//展示用法
func PrintUsage() {
	fmt.Println("Usage:")
	fmt.Printf("\tcreatewallet --创建钱包. \n")
	fmt.Printf("\tgetaddresslist --获取钱包地址列表. \n")
	fmt.Printf("\tcreateblockchain  -address address -- 地址.\n")
	fmt.Printf("\taddblock -data  DATA  -- 交易数据.\n")
	fmt.Printf("\tprintblockchain -- 输出区块链的信息.\n")
	fmt.Printf("\tsend -from addr -to addr -amout AMOUNT -- 转账.\n")
	fmt.Printf("\tgetbalance -from addr  -- 查询余额.\n")
}

//校验，如果只输入了程序命令，就输出指令用法并且退出程序
func IsValidArgs() {
	if len(os.Args) < 2 {
		PrintUsage()
		os.Exit(1)
	}
}
////添加区块
//func (cli *CLI) addBlock(txs []*Transaction) {
//	if dbExists() == false {
//		log.Info("数据库不存在")
//		os.Exit(1)
//	}
//	blockchain := BlockchainObject()
//	defer blockchain.DB.Close()
//	blockchain.AddBlock(txs)
//}
//运行函数
func (cli *CLI) Run() {
	//1. 检测参数数量
	IsValidArgs()

	//2. 新建命令
	createWalletsCmd := flag.NewFlagSet("createwallet", flag.ExitOnError)
	getAddressListsCmd := flag.NewFlagSet("getaddresslist", flag.ExitOnError)
	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	createBlcWithGenesisCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)

	//3. 获取命令行参数
	//flagAddBlockArg := addBlockCmd.String("data", "send 100 BTC to everyone", "交易数据...")
	flagCreateBlcWithGenesisArg := createBlcWithGenesisCmd.String("address", "", "地址...")
	flagFromArg := sendCmd.String("from", "", "发送者")
	flagToArg := sendCmd.String("to", "", "接收者")
	flagAmountArg := sendCmd.String("amount", "", "转账金额")
	flagBalanceArg := getBalanceCmd.String("from", "", "转账金额")
	//flagPrintChainArg:=printChainCmd.String("data","send 100 BTC to everyone","交易数据")
	//flagcreateBlcWithGenesisArg:=createBlcWithGenesisCmd.String("data","send 100 BTC to everyone","交易数据")
	switch os.Args[1] {

	case "createwallet":
		err := createWalletsCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panicf("parse cmd of createwallet failed! %v\n", err.Error())
		}
	case "getaddresslist":
		err := getAddressListsCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panicf("parse cmd of getaddresslist failed! %v\n", err.Error())
		}
	case "getbalance":
		err:=getBalanceCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panicf("parse cmd of getbalance failed! %v\n", err.Error())
		}
	case "send":
		err := sendCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panicf("parse cmd of send failed! %v\n", err.Error())
		}
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])

		if err != nil {
			log.Panicf("parse cmd of addblock failed! %v\n", err.Error())
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panicf("parse cmd of printchain failed! %v\n", err.Error())
		}
	case "createblockchain":
		err := createBlcWithGenesisCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panicf("parse cmd of createblockchain failed! %v\n", err.Error())
		}
	default:
		PrintUsage()
		os.Exit(1)
	}

	//添加查询余额命令
	if getBalanceCmd.Parsed(){
		if *flagBalanceArg==""{
			PrintUsage()
			os.Exit(1)
		}

		cli.getBalance(*flagBalanceArg)
	}

	//添加转账命令
	if sendCmd.Parsed() {
		if *flagFromArg == "" || *flagToArg == "" || *flagAmountArg == "" {
			PrintUsage()
			os.Exit(1)
		}

		fmt.Printf("\tFROM:[%s]\n", utils.JSONToArray(*flagFromArg))
		fmt.Printf("\tTO:[%s]\n", utils.JSONToArray(*flagToArg))
		fmt.Printf("\tAmount:[%s]\n", utils.JSONToArray(*flagAmountArg))

		cli.send(utils.JSONToArray(*flagFromArg), utils.JSONToArray(*flagToArg), utils.JSONToArray(*flagAmountArg))
	}

	//if addBlockCmd.Parsed() {
	//	if *flagAddBlockArg == "" {
	//		PrintUsage()
	//		os.Exit(1)
	//	}
	//
	//	cli.addBlock([]*Transaction{})
	//}

	if printChainCmd.Parsed() {
		cli.printchain()
	}

	if createBlcWithGenesisCmd.Parsed() {
		if *flagCreateBlcWithGenesisArg == "" {
			PrintUsage()
			os.Exit(1)
		}
		cli.createBlockchinWithGenesis(*flagCreateBlcWithGenesisArg)
	}

	if createWalletsCmd.Parsed(){
		cli.CreateWallets()
	}

	if getAddressListsCmd.Parsed(){
		cli.getAddressLists()
	}

}
