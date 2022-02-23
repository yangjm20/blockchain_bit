package BLC

import (
	"github.com/boltdb/bolt"
	"github.com/fabric_assert/blockchain_bit/pkg/log"
)

//创建迭代器对象
func (bc *BlockChain) Iterator() *BlockChainIterator {
	return &BlockChainIterator{bc.DB, bc.Tip}
}

//遍历
func (bcit *BlockChainIterator) Next() *Block {
	var block *Block
	err:=bcit.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BlockTableName))
		if nil != b {
			blockBytes := b.Get(bcit.CurrentHash)
			block=DeserializeBlock(blockBytes)
			//更新迭代器中当前区块的hash值
			bcit.CurrentHash=block.PreBlockHash
		}
		return nil
	})

	if nil!=err{
		log.Panicf("iterator the db of blockchain failed ! %v",err.Error())
	}

	return block
}