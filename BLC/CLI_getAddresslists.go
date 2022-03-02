package BLC

import (
	"fmt"
	"github.com/fabric_assert/blockchain_bit/wallets"
)

func (cli *CLI)getAddressLists()  {
	fmt.Println("打印所有钱包地址的集合")
	wallets,_:=wallets.NewWallets()
	//if err!=nil{
	//	log.Panicf("new wallet is failed! %v\n",err.Error())
	//}
	for address,wallet :=range wallets.Wallets{
		fmt.Printf("address:[%s],wallet:[%v]\n",address,wallet)
	}
}
