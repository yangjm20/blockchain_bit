package utils

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
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

func JSONToArray(jsonString string) []string  {
	var strArr []string
	if err:=json.Unmarshal([]byte(jsonString),&strArr);err!=nil{
		log.Panicf("json to []string failed! %v\n",err.Error())
	}
	return strArr
}

//反转切片
func Reverse(data []byte)  {
	for i,j:=0,len(data)-1;i<j;i,j=i+1,j-1{
		data[i],data[j]=data[j],data[i]
	}
}