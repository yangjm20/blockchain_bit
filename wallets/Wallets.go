package wallets

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"github.com/fabric_assert/blockchain_bit/pkg/log"
	"io/ioutil"
	"os"
)

//钱包集合的文件
const walletFile = "wallets.dat"
//钱包的集合
type Wallets struct {
	Wallets map[string]*Wallet
}
//创建一个钱包集合
func NewWallets() (*Wallets ,error) {
	//1. 判断文件是否存证
	if _,err:=os.Stat(walletFile);os.IsNotExist(err){
		wallets:=&Wallets{}
		wallets.Wallets=make(map[string]*Wallet)
		return wallets,err
	}

	//2. 如果存在，读取内容
	fileContent,err:=ioutil.ReadFile(walletFile)
	if nil!=err{
		log.Panicf("get file content failed! %v\n",err.Error())
	}
	var wallets Wallets
	//register适用于需要解析的参数中包含interface
	gob.Register(elliptic.P256())
	decode:=gob.NewDecoder(bytes.NewReader(fileContent))
	err=decode.Decode(&wallets)
	if nil!=err{
		log.Panicf("decode file content failed! %v\n",err.Error())
	}
	return &wallets,nil
}

//创建新的钱包,并将其添加到钱包集合中
func (w *Wallets)CreateWallet()  {
	wallet:=NewWallet()
	address:=wallet.GetAddress()
	w.Wallets[string(address)]=wallet

	//把钱包存储到文件中
	w.SaveWallets()
}

//持久化钱包信息(写入文件)
func (w *Wallets)SaveWallets()  {
	var content bytes.Buffer
	gob.Register(elliptic.P256())
	encode:=gob.NewEncoder(&content)
	//序列化钱包数据
	err:=encode.Encode(&w)
	if err!=nil{
		log.Panicf("encode the struct of wallets failed!\n",err.Error())
	}
	//清空文件，再去存储（此处只保存了一条数据）
	err=ioutil.WriteFile(walletFile,content.Bytes(),0644)
	if err!=nil{
		log.Panicf("write the content of wallets to file [%s] failed!\n",walletFile,err.Error())
	}
}