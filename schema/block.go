package schema

type Block struct {
	PreviousHash *string `json:"previousHash"`
	Data         string  `json:"data"`
	Hash         string  `json:"hash"`
}
