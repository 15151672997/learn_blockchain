package main

import (
	"fmt"
	"publicChain/part3-Basic-Prototype/BLC"
)

func main() {

	genesisBlockchain := BLC.CreateBlockchainWithGenesisBlock()

	fmt.Println(*genesisBlockchain)
	fmt.Println(genesisBlockchain.Blocks)
}
