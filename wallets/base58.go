package wallets

import (
	"bytes"
	"fmt"
	"github.com/fabric_assert/blockchain_bit/pkg/utils"
	"math/big"
)

//var baseAlphabet = []byte("123456789" +
//	"abcdefghijkmnopqrstuvwxyz" +
//	"ABCDEFGHJKLMNPQRSTUVWXYZ")

var baseAlphabet = []byte("123456789" +
	"ABCDEFGHJKLMNPQRSTUVWXYZ"+
	"abcdefghijkmnopqrstuvwxyz")

func Base58Encode(input []byte) []byte {
	var result []byte
	x := big.NewInt(0).SetBytes(input) //bytes 转换为bigint
	fmt.Printf("x:%v\n", x)
	base := big.NewInt(int64(len(baseAlphabet))) //设置一个base58求模的基数
	zero := big.NewInt(0)
	mod := &big.Int{}  //余数
	for x.Cmp(zero) != 0 {
		x.DivMod(x, base, mod) //求余
		//以余数为下标取值
		result = append(result, baseAlphabet[mod.Int64()])
	}

	//反转切片
	utils.Reverse(result)
	for b := range input { //b代表切片下标
		if b == 0x00 {
			result = append([]byte{baseAlphabet[0]}, result...)
		} else {
			break
		}
	}
	fmt.Printf("result:%s\n",result)
	return result
}

func Base58Decode(input []byte) []byte {
	result:=big.NewInt(0)
	zeroBytes:=0
	for b:= range input{
		if b==0x00{
			zeroBytes++
		}
	}
	data:=input[zeroBytes:]
	for _,b:=range data{
		charIndex:=bytes.IndexByte(baseAlphabet,b)
		result.Mul(result,big.NewInt(58))
		result.Add(result,big.NewInt(int64(charIndex)))
	}

	decoded:=result.Bytes()
	decoded=append(bytes.Repeat([]byte{byte(0x00)},zeroBytes),decoded...)
	return decoded
}