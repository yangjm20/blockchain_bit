package BLC

type UTXO struct {
	//UTXO 所对应的交易哈希
	TxHash []byte
	//UTXO在其所属交易中的索引
	Index  int
	//OUTPUT
	Output *TxOutput
}
