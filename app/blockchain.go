package app

import (
	"fmt"
	"solyn/reth/helpers"
	"solyn/reth/schema"
)

func InitBlock() {

	blockchain := helpers.GetBlockchain()
	fmt.Println(blockchain)
	newBlock := schema.Block{
		PreviousHash: "109291code",
		Data:         "This is data12",
		Hash:         "hashh31",
	}
	// push newblock to blockchain
	blockchain = append(blockchain, newBlock)
	fmt.Println(blockchain)
	helpers.WriteBlockchain(blockchain)
}
