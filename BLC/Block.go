package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"github.com/fabric_assert/blockchain_bit/pkg/utils"
	"log"

	"time"
)

type Block struct {
	TimeStamp    int64 //区块时间戳
	Heigth       int64 //区块高度
	PreBlockHash []byte
	Hash         []byte
	Data         []byte
	Nonce        int64 //用来生成工作量证明的hash
}

func NewBlock(height int64, preBlockHash []byte, data []byte) *Block {
	var block Block
	block = Block{Heigth: height, PreBlockHash: preBlockHash, Data: data, TimeStamp: time.Now().Unix()}
	//block.SetHash()
	pow := NewProofOfWork(&block)
	hash, nonce := pow.Run()
	block.Hash = hash
	block.Nonce = nonce
	return &block
}

func (b *Block) SetHash() {
	timeStampBytes := utils.IntToHex(b.TimeStamp)
	heightBytes := utils.IntToHex(b.Heigth)
	blockBytes := bytes.Join([][]byte{heightBytes, timeStampBytes, b.PreBlockHash, b.Data}, []byte{})
	hash := sha256.Sum256(blockBytes)
	b.Hash = hash[:]
}

func CreateGenesisBlock(data string) *Block {
	block := NewBlock(1, nil, []byte(data))

	return block
}

//序列化，将区块结构序列化为[]byte
func (block *Block) Serialize() []byte{
	var result bytes.Buffer
	encoder:=gob.NewEncoder(&result) //新建encode对象
	if err:=encoder.Encode(block);nil!=err{
		log.Panicf("serialize the block to byte failed ! %v",err.Error())
	}

	return result.Bytes()
}

//反序列化，将字节数组结构反序列化为区块结构
func DeserializeBlock(blockBytes []byte) *Block{
	var block *Block
	block = &Block{}
	decoder:=gob.NewDecoder(bytes.NewReader(blockBytes))

	if err:=decoder.Decode(block);nil!=err{
		log.Panicf("Deserialize the []byte to Bolck! %v",err.Error())
	}
	return block
}
