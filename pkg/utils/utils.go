package utils

import (
	"bytes"
	"encoding/binary"
	"github.com/fabric_assert/blockchain_bit/pkg/log"
)

func IntToHex(data int64) []byte  {
	buffer:=new(bytes.Buffer)
	err:=binary.Write(buffer,binary.BigEndian,data)
	if nil!=err{
		log.Panicf("int to []byte failed! %v\n",err.Error())
	}
	return buffer.Bytes()
}

