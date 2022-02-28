package BLC


//交易输出
type TxOutput struct {
	Value int64 //金额
	ScriptPubkey string//钱是谁的，账户
}

//output 身份验证

func (out *TxOutput) UnLockScriptPubkeyWithAddress(address string) bool {
	return  out.ScriptPubkey== address
}