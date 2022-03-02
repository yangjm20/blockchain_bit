package wallets

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"golang.org/x/crypto/ripemd160"
	"log"
)

const version = byte(0x00)

//checksum 长度
const addressChecksumLen = 4

type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublickKey []byte
}

func NewWallet() *Wallet {
	priv, pubKey := newKeyPair()
	return &Wallet{priv, pubKey}
}

func newKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()
	priv, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panicf("ecdsa generate key failed! %v\n", err.Error())
	}

	pubKey := append(priv.PublicKey.X.Bytes(), priv.PublicKey.Y.Bytes()...)

	return *priv, pubKey
}

//对公钥匙进行双哈希（sha256->ripemd160）
func Ripemd160Hash(pubKey []byte) []byte {
	//sha256
	hash := sha256.New()
	hash.Write(pubKey)
	hashed := hash.Sum(nil)

	ripeHash := ripemd160.New()
	ripeHash.Write(hashed)
	ripeHashed := ripeHash.Sum(nil)

	return ripeHashed
}

//通过钱包获取地址
func (w *Wallet)GetAddress() []byte {
	//1 获取160hash
	ripe160Hash:=Ripemd160Hash(w.PublickKey)
	//2 生成version并加入hash中
	version_ripemd160Hash:=append([]byte{version},ripe160Hash...)
	//3.生成校验和数据
	checksumBytes:=GenerateCheckSum(version_ripemd160Hash)
	//4.拼接校验和
	bytes:=append(version_ripemd160Hash,checksumBytes...)

	//5 调用base58Encode 生成地址
	base58Bytes:=Base58Encode(bytes)
	return base58Bytes
}

func GenerateCheckSum(payload []byte)[]byte  {
	hash_first:=sha256.Sum256(payload)
	hash_second:=sha256.Sum256(hash_first[:])
	return hash_second[:addressChecksumLen]
}

//判断地址有效性
func IsValidForAddress(address []byte)bool  {
	//1地址通过base58decode
	version_pubkey_checksum:=Base58Decode(address) //25位
	//2拆开，进行校验和的校验
	checkSumBytes:=version_pubkey_checksum[len(version_pubkey_checksum)-addressChecksumLen:]
	version_ripemd160:=version_pubkey_checksum[:len(version_pubkey_checksum)-addressChecksumLen]
	checkBytes:=GenerateCheckSum(version_ripemd160)
	if bytes.Compare(checkSumBytes,checkBytes)==0{
		return true
	}
	return false

}
