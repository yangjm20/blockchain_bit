package BLC

import (
	"bytes"
	"crypto/sha256"
	"github.com/fabric_assert/blockchain_bit/pkg/log"
	"github.com/fabric_assert/blockchain_bit/pkg/utils"
	"math/big"
)

//目标难度值，代表生成的哈希值targetBit 位为0，才能满足条件
const targetBit = 8

//工作量证明
type ProofOfWork struct {
	Block  *Block
	target *big.Int
}

//创建新的POW对象
func NewProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)

	target = target.Lsh(target, 256-targetBit)

	//hash 256位
	//前16位都为0
	//左移

	return &ProofOfWork{block, target}
}

//开始工作量证明
func (pow *ProofOfWork) Run() ([]byte, int64) {
	//1.数据拼接
	var nonce = 0  //碰撞次数
	var hash [32]byte  //hash值
	var hashInt big.Int

	for {
		dataBytes := pow.prepareData(nonce)
		hash = sha256.Sum256(dataBytes)
		hashInt.SetBytes(hash[:])
		//log.Infof("hash : \r%x", hash)
		//难度比较
		if pow.target.Cmp(&hashInt) == 1 {
			break
		}

		nonce++
	}
	log.Infof("hash : \r%x", hash)
	log.Infof("\n碰撞次数：%d\n",nonce)
	return hash[:], int64(nonce)
}

//准备数据，将区块相关的属性拼接起来，返回一个字节数组
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join([][]byte{
		pow.Block.PreBlockHash,
		pow.Block.HashTransactions(),
		utils.IntToHex(pow.Block.Heigth),
		utils.IntToHex(pow.Block.TimeStamp),
		utils.IntToHex(int64(nonce)),
		utils.IntToHex(targetBit),
	}, []byte{})
	return data
}
