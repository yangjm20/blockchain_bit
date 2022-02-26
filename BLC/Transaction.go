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
func NewSimpleTransaction(from string,to string ,amount int)  *Transaction{
	var txInputs []*TxInput
	var txOutputs []*TxOutput

	//消费
	txInput:=&TxInput{[]byte("8df099565e4bb780b86082988e363841916fed7528738e02750708914f1c3c69"),0,from}
	txInputs=append(txInputs,txInput)

	txOutPut:=&TxOutput{int64(amount),to}
	txOutputs=append(txOutputs,txOutPut)
	txOutPut=&TxOutput{10-int64(amount),from}
	txOutputs=append(txOutputs,txOutPut)

	tx:=&Transaction{nil,txInputs,txOutputs}
	tx.HashTransaction()

	return tx
}
