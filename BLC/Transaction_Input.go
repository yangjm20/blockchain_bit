package BLC

//交易输入
type TxInput struct {
	TxHash    []byte //交易哈希（不是当前交易的hash）
	Vout      int    //引用上一笔交易的output索引
	ScriptSig string //用户名
}