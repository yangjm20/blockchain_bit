package BLC

import (
	"github.com/fabric_assert/blockchain_bit/pkg/log"
	"os"
)

func (cli *CLI) send(from, to, amount []string) {
	if dbExists() == false {
		log.Info("数据库不存在")
		os.Exit(1)
	}
	blockchain := BlockchainObject() //获取区块链对象
	defer blockchain.DB.Close()
	blockchain.MineNewBlock(from, to, amount)
}
