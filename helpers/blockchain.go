package helpers

import (
	"encoding/json"
	"log"
	"os"
	"solyn/reth/schema"
)

func GetBlockchain() []schema.Block {
	var blockchain []schema.Block
	file, err := os.ReadFile("blockchain.json") // For read access.
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal([]byte(file), &blockchain)
	return blockchain
}

func WriteBlockchain(blockchain []schema.Block) {
	file, err := json.MarshalIndent(blockchain, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	_ = os.WriteFile("blockchain.json", file, 0644)
}
