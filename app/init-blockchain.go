package app

import (
	"math/rand"
	"solyn/reth/helpers"
	"solyn/reth/schema"
	"strconv"
)

func InitGenesisBlock() {
	currentBlockchain, _ := helpers.GetBlockchain()
	if len(currentBlockchain) > 0 {
		return
	}
	var genesisBlock = schema.Block{
		PreviousHash: nil,
		Data:         "GenesisBlock",
		Hash:         strconv.Itoa(rand.Intn(1000)),
	}

	blockchain := []schema.Block{genesisBlock}
	helpers.WriteBlockchain(blockchain)
}
