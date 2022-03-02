package BLC

//交易输入
type TxInput struct {
	TxHash    []byte //交易哈希（不是当前交易的hash）
	Vout      int    //引用上一笔交易的output索引
	ScriptSig string //用户名
}

//判断能不能引用指定地址的OUTPUT
func (in *TxInput) UnLockWithAddress(address string) bool {
	return in.ScriptSig == address
}
