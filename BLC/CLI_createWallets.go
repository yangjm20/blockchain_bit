package BLC

import (
	"fmt"
	"github.com/fabric_assert/blockchain_bit/wallets"
)

//创建钱包集合
func (cli *CLI)CreateWallets() {
	wallets,_:=wallets.NewWallets()
	//if err!=nil{
	//	log.Panicf("new wallet is failed! %v\n",err.Error())
	//}
	wallets.CreateWallet()

	fmt.Printf("wallets:%v\n",wallets)

}

