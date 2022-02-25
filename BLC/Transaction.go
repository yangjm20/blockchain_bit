package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
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
