package main

import (
	"github.com/fabric_assert/blockchain_bit/BLC"
	"github.com/fabric_assert/blockchain_bit/pkg/log"
)

func main() {
	// logger配置
	opts := &log.Options{
		Level:            "debug",
		Format:           "console",
		EnableColor:      true, // if you need output to local path, with EnableColor must be false.
		DisableCaller:    false,
		OutputPaths:      []string{"test.log", "stdout"},
		ErrorOutputPaths: []string{"error.log"},
	}
	// 初始化全局logger
	log.Init(opts)
	defer log.Flush()



	//block:=BLC.NewBlock(1,nil,[]byte("this is the base concept block"))
	//block:=BLC.CreateGenesisBlock("this is the basic concept block")
	//bc:=BLC.CreateBlockChainWithGenesisBlock()
	//log.Infof("block : %v \n",bc)
	//bc.AddBlock(bc.Blocks[len(bc.Blocks)-1].Heigth+1,[]byte("Aii send 100 BTC to Bob"),bc.Blocks[len(bc.Blocks)-1].PreBlockHash)
	//bc.AddBlock(bc.Blocks[len(bc.Blocks)-1].Heigth+1,[]byte("Bob send 100 BTC to Tommy"),bc.Blocks[len(bc.Blocks)-1].PreBlockHash)

	//for i:=0;i<len(bc.Blocks);i++{
	//	log.Infof("the %d'th block is %v \n",i,bc.Blocks[i])
	//}


	blockchain:=BLC.CreateBlockChainWithGenesisBlock()
	//
	//
	//db:=blockchain.DB
	//
	//blockchain.AddBlock([]byte("Send 100 btc to troy"))
	//blockchain.AddBlock([]byte("Send 200 btc to Alice"))
	//blockchain.AddBlock([]byte("Send 50 btc to Blob"))
	//
	//blockchain.PrintChain()
	//
	//defer db.Close()
	//BLC.PrintUsage()
	cli:=BLC.CLI{blockchain}
	cli.Run()




}
