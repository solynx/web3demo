package helpers

import (
	"encoding/json"
	"log"
	"os"
	"solyn/reth/schema"
)

func GetBlockchain() ([]schema.Block, error) {
	var blockchain []schema.Block
	file, err := os.ReadFile("blockchain.json") // For read access.
	if err != nil {
		return blockchain, err
	}
	json.Unmarshal([]byte(file), &blockchain)
	return blockchain, nil
}

func WriteBlockchain(blockchain []schema.Block) {
	file, err := json.MarshalIndent(blockchain, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	_ = os.WriteFile("blockchain.json", file, 0644)
}
