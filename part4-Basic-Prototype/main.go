package main

import (
	"fmt"
	"publicChain/part3-Basic-Prototype/BLC"
)

func main() {

	blockchain := BLC.CreateBlockchainWithGenesisBlock()

	fmt.Println(*blockchain)
	fmt.Println(blockchain.Blocks)

	//添加新的区块
	blockchain.AddBlockToBlockchain("send 100 RMB to zhangqiang", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
}
