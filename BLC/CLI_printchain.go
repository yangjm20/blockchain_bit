package BLC

import (
	"github.com/fabric_assert/blockchain_bit/pkg/log"
	"os"
)

//输出区块链信息

func (cli *CLI) printchain() {
	if dbExists() == false {

		log.Info("数据库不存在")

		os.Exit(1)
	}
	blockchain := BlockchainObject()
	defer blockchain.DB.Close()
	blockchain.PrintChain()
}
