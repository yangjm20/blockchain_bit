package BLC

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/fabric_assert/blockchain_bit/pkg/log"
	"math/big"
	"os"
	"strconv"
)

const DbName = "bc.db"          //存储区块数据的数据库文件
const BlockTableName = "blocks" //表（桶）名称

//区块链基本结构
type BlockChain struct {
	DB  *bolt.DB // 数据库
	Tip []byte   //最新区块的哈希值
}

//区块链迭代器结构
type BlockChainIterator struct {
	DB          *bolt.DB // 数据库
	CurrentHash []byte   //当前区块的哈希值
}

func dbExists() bool {
	if _,err:=os.Stat(DbName);os.IsNotExist(err){
		return false
	}
	return true
}

//初始化区块链
func CreateBlockChainWithGenesisBlock(address string) *BlockChain {

	if dbExists(){
		log.Info("创世区块已经存在...")
		os.Exit(1)

	}

	db, err := bolt.Open(DbName, 0600, nil)
	if err != nil {
		log.Panicf("open the db failed : %v\n", err.Error())
	}

	var blockHash []byte //需要存储到数据库中到区块hash

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BlockTableName))
		if nil == b {
			//添加创世区块
			b, err = tx.CreateBucket([]byte(BlockTableName))
			if nil != err {
				log.Panicf("create the bucket [%s] failed %v \n", BlockTableName, err.Error())
			}
		}
		if nil != b {

			//生成创世区块链
			tx:=NewCoinbaseTransaction(address)
			genesisBlock := CreateGenesisBlock([]*Transaction{tx})
			err := b.Put(genesisBlock.Hash, genesisBlock.Serialize())
			if nil != err {
				log.Panicf("put the data of genesisBlock to db failed %v\n", err.Error())
			}

			//存储最新区块的哈希值
			err = b.Put([]byte("1"), genesisBlock.Hash)
			if nil != err {
				log.Panicf("put the hash of latest block to db failed ! %v", err.Error())
			}
			blockHash = genesisBlock.Hash
		}
		return nil
	})
	if nil != err {
		log.Panicf("update the block to db failed ! %v", err.Error())
	}

	return &BlockChain{db, blockHash}
}

func (bc *BlockChain) AddBlock(txs []*Transaction) {
	//newBlock := NewBlock(height, preBlockHash, data)
	//bc.Blocks = append(bc.Blocks, newBlock)

	//更新数据
	err := bc.DB.Update(func(tx *bolt.Tx) error {
		//1 获取数据表
		b := tx.Bucket([]byte(BlockTableName))
		if nil != b { //2. 确保表存在

			//3.获取最新区块的哈希
			//newLastHash := b.Get([]byte("1"))
			blockBytes := b.Get(bc.Tip)
			latest_block := DeserializeBlock(blockBytes)
			//4.创建新区块


			newblock := NewBlock(latest_block.Heigth+1, latest_block.Hash, txs)
			err := b.Put(newblock.Hash, newblock.Serialize())
			if err != nil {
				log.Panicf("put the data of new block into db failed! %v \n", err.Error())
			}

			err = b.Put([]byte("1"), newblock.Hash)
			if nil != err {
				log.Panicf("put the hash of the newest block into db failed ! %v", err.Error())
			}

			bc.Tip = newblock.Hash

		}

		return nil
	})

	if nil != err {
		log.Panicf("update the db of the block failed ! %v", err.Error())
	}

}

//遍历输出所有区块的信息
func (bc *BlockChain) PrintChain() {
	var block *Block
	bcit:=bc.Iterator()

	for {
		fmt.Println("--------------------------------------")
		block=bcit.Next()
		fmt.Printf("\tHeight:%d\n",block.Heigth)
		fmt.Printf("\tTimeStamp:%d\n",block.TimeStamp)
		fmt.Printf("\tPreHash:%x\n",block.PreBlockHash)
		fmt.Printf("\tHash:%x\n",block.Hash)
		fmt.Printf("\ttTransaction:%v\n",block.Txs)
		fmt.Println("\t{")

		for _,tx:= range block.Txs{
			fmt.Printf("\t\ttx-hash:%x \n",tx.TxHash)
			fmt.Println("\t\t输入...")
			for _,vin:= range tx.Vins{
				fmt.Printf("\t\t{input{%x,%v,%v}\n",vin.TxHash,vin.Vout,vin.ScriptSig)
				//fmt.Printf("\t\t vin-tx-hash:%x \n",vin.TxHash)
				//fmt.Printf("\t\t vout:%v \n",vin.Vout)
				//fmt.Printf("\t\t scriptSig:%v \n",vin.ScriptSig)
			}
			fmt.Println("\t\t输出...")
			for _,out:= range tx.Vout{
				fmt.Printf("\t\toutput{%v,%v}\n",out.Value,out.ScriptPubkey)
				//fmt.Printf("\t\t vout-value:%v \n",out.Value)
				//fmt.Printf("\t\t vout:%v \n",out.ScriptPubkey)
			}

		}
		fmt.Println("\t}")


		var hashInt big.Int
		hashInt.SetBytes(block.PreBlockHash)
		if big.NewInt(0).Cmp(&hashInt) == 0{
			break
		}

	}

}

func BlockchainObject()*BlockChain{
	db,err:=bolt.Open(DbName,0600,nil)
	if nil!=err{
		log.Panicf("get the object of blockchain failed! %v \n",err.Error())
	}

	var tip []byte
	err=db.View(func(tx *bolt.Tx) error {
		b:=tx.Bucket([]byte(BlockTableName))
		if b!=nil{
			tip=b.Get([]byte("1")) //最新区块的hash值
			return nil
		}
		return errors.New("数据库无数据")
	})
	if err!=nil{
		log.Panicf("the database is null %v ",err.Error())

	}

	return &BlockChain{db,tip}

}

//挖矿打通过接收交易，进行打包确认，最终生成新的区块
//t
func (bc *BlockChain)MineNewBlock(from ,to ,amount []string)  {
	//接收交易
	var txs []*Transaction //要打包的交易
	//打包交易
	value,_:=strconv.Atoi(amount[0])
	tx:=NewSimpleTransaction(from[0],to[0],value,bc)
	txs=append(txs,tx)
	//生成新的区块
	var block *Block
	//从数据中获取最新的区块
	bc.DB.View(func(tx *bolt.Tx) error {
		b:=tx.Bucket([]byte(BlockTableName))
		if nil!=b{
			hash:=b.Get([]byte("1"))
			blockBytes:=b.Get(hash)
			block=DeserializeBlock(blockBytes)

			return nil
		}

		return errors.New("数据库表中无数据")
	})
	block=NewBlock(block.Heigth+1,block.Hash,txs)

	//持久化新的区块
	bc.DB.Update(func(tx *bolt.Tx) error {
		b:=tx.Bucket([]byte(BlockTableName))
		if nil!=b{

			err:=b.Put(block.Hash,block.Serialize())
			if err != nil {
				log.Panicf("put the data of new block into db failed! %v \n", err.Error())
			}

			err = b.Put([]byte("1"), block.Hash)
			if nil != err {
				log.Panicf("put the hash of the newest block into db failed ! %v", err.Error())
			}

			bc.Tip=block.Hash
			return nil
		}
		return errors.New("数据库表不存在")
	})
}

func (bc *BlockChain)UnUTXOS(address string) []*UTXO {
	var unUTXOS []*UTXO
	bcit:=bc.Iterator()
	//存储所有已花费的输出
	//key :每个input所引用的交易哈希
	//value:output 索引列表
	spentTxOutputs:=make(map[string][]int)

	for  {
		block:=bcit.Next() //获取每一个区块信息

		for _,tx:= range block.Txs{ //遍历每个区块中的交易

			//先查找输入
			if !tx.IsCoinbaseTransaction(){

				for _,in:= range tx.Vins{
					//验证地址
					if in.UnLockWithAddress(address){
						//添加到已花费输入中
						key:=hex.EncodeToString(in.TxHash)
						spentTxOutputs[key]=append(spentTxOutputs[key],in.Vout)
					}
				}
			}


			work:
			for index,vout:=range tx.Vout{ //再查找vout

				//地址验证（检查输出是否属于传入地址）
				if vout.UnLockScriptPubkeyWithAddress(address){
					//判断output是否是一个未花费的输出
					//判断已花费输出是否为空
					if len(spentTxOutputs)!=0{
						var isSpentTxOutput bool

						for txHash,indexArrary:=range spentTxOutputs{
							for _,i:=range indexArrary{
								if txHash==hex.EncodeToString(tx.TxHash) && index==i{
									isSpentTxOutput =true
									continue work
								}
							}
						}
						if isSpentTxOutput ==false{
							utxo:=&UTXO{tx.TxHash,index,vout}
							unUTXOS = append(unUTXOS,utxo)
						}
					}else{
						fmt.Println("已花费输出为空")
						utxo:=&UTXO{tx.TxHash,index,vout}
						unUTXOS=append(unUTXOS,utxo)
					}
				}
			}

		}


		var hashInt big.Int
		hashInt.SetBytes(block.PreBlockHash)
		if big.NewInt(0).Cmp(&hashInt) == 0{
			break
		}

	}

	fmt.Printf("the address is %s\n",address)
	return unUTXOS
}

func (bc *BlockChain)getBalance(address string)int64 {
	utxos:=bc.UnUTXOS(address)
	var amount int64
	for _,utxo:=range utxos{
		amount+=utxo.Output.Value
	}
	return amount
}

//转账时候查找可用的UTXO
//查找可用的UTXO（遍历），超过需要的资金即可中断
func (bc *BlockChain) FindSpendableUTXO(from string,amount int64) (int64,map[string][]int){
	//查找出来的UTXO的值总和
	var value int64
	//可用的UTXO
	spendableUTXO:=make(map[string][]int)
	//获取所有的UTXO
	utxos:=bc.UnUTXOS(from)
	//遍历
	for _,utxo:=range utxos{
		value+=utxo.Output.Value
		hash:=hex.EncodeToString(utxo.TxHash)
		spendableUTXO[hash] = append(spendableUTXO[hash],utxo.Index)
		if value>=amount{
			break
		}
	}
	if value<amount{
		fmt.Printf("%s 余额不足",from)
		os.Exit(1)
	}
	return value,spendableUTXO
}