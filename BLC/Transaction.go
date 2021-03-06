package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"github.com/fabric_assert/blockchain_bit/pkg/log"
)

type Transaction struct {
	TxHash []byte

	//输入
	Vins []*TxInput
	//输出
	Vout []*TxOutput
}

//生成交易哈希
func (tx *Transaction)HashTransaction() {
	var result bytes.Buffer
	encode := gob.NewEncoder(&result)
	err:=encode.Encode(tx)
	if err!=nil{
		log.Panicf("tx hash generate failed! %v",err.Error())
	}

	hash:=sha256.Sum256(result.Bytes())
	tx.TxHash = hash[:]

}



//生成coinbase交易
func NewCoinbaseTransaction(address string) *Transaction{
	//输入
	txInput:=&TxInput{[]byte{},-1,"Genesis Data"}
	//输出
	txOutput:=&TxOutput{10,address}
	//hash
	txCoinbase:=&Transaction{nil,[]*TxInput{txInput},[]*TxOutput{txOutput}}

	txCoinbase.HashTransaction()
	return txCoinbase

}

//生成转账交易
func NewSimpleTransaction(from string,to string ,amount int,blockchain *BlockChain,txs []*Transaction)  *Transaction{
	var txInputs []*TxInput
	var txOutputs []*TxOutput

	//查找指定地址的可用UTXO
	money,spendableUTXODic:=blockchain.FindSpendableUTXO(from,int64(amount),txs)
	fmt.Printf("money:%v\n",money)

	for txHash,indexArray:=range spendableUTXODic{
		txHashBytes,_:=hex.DecodeString(txHash)
		for _,index:=range indexArray{
			//此处的输出是需要消费的，必然会被其他的交易的输入所引用
			//消费
			txInput:=&TxInput{txHashBytes,index,from}
			txInputs=append(txInputs,txInput)
		}
	}

	//转账
	txOutPut:=&TxOutput{int64(amount),to}
	txOutputs=append(txOutputs,txOutPut)
	//找零
	txOutPut=&TxOutput{money-int64(amount),from}
	txOutputs=append(txOutputs,txOutPut)

	//生成交易
	tx:=&Transaction{nil,txInputs,txOutputs}
	tx.HashTransaction()

	return tx
}


//判断指定交易是否是一个coinbase交易
func (tx *Transaction)IsCoinbaseTransaction() bool {
	//fmt.Printf("iscoinbase txhash:%s,vout=%d",tx.Vins[0].TxHash,tx.Vins[0].Vout)
	return len(tx.Vins[0].TxHash)==0 && tx.Vins[0].Vout==-1
}
