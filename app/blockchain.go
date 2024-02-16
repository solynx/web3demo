package app

import (
	"fmt"
	"math/rand"
	"solyn/reth/helpers"
	"solyn/reth/schema"
	"strconv"
)

func InitBlock() {

	blockchain, _ := helpers.GetBlockchain()
	fmt.Println(blockchain)
	newBlock := schema.Block{
		PreviousHash: &blockchain[len(blockchain)-1].Hash,
		Data:         "This is data12",
		Hash:         strconv.Itoa(rand.Intn(1000)),
	}
	// push newblock to blockchain
	blockchain = append(blockchain, newBlock)
	fmt.Println(blockchain)
	helpers.WriteBlockchain(blockchain)
}
