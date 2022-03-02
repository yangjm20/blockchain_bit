package BLC

import "fmt"

//查询余额
func (cli *CLI)getBalance(from string) {
	blockchain := BlockchainObject() //获取区块链对象
	defer blockchain.DB.Close()
	amount:=blockchain.getBalance(from)

	fmt.Printf("\t地址：%s的余额为%d\n",from,amount)
}
