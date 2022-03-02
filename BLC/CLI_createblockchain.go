package BLC
//创建创世区块
func (cli *CLI) createBlockchinWithGenesis(address string) {
	CreateBlockChainWithGenesisBlock(address)
}
