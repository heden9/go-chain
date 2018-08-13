package core

import (
	"time"
	"encoding/hex"
	"crypto/sha256"
)

type Block struct {
	Index int64 // 区块编号
	Timestamp int64 // 区块时间戳
	PrevBlockHash string // 上一个区块hash
	Hash string // 当前区块hash

	Data string // 区块数据
}

func GetBlockHash(b *Block) string {
	var (
		blockData string
		hashInBytes [32]byte
	)

	blockData = string(b.Index + b.Timestamp) + b.PrevBlockHash + b.Data
	hashInBytes = sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(hashInBytes[:])
}

func CreateNewBlock(prevBlock *Block, data string) *Block {
	var (
		newBlock *Block
	)
	newBlock = &Block{
		Index: prevBlock.Index + 1,
		PrevBlockHash: prevBlock.Hash,
		Timestamp: time.Now().Unix(),
		Data: data,
	}
	newBlock.Hash = GetBlockHash(newBlock)

	return newBlock
}

func CreateGenesisBlock() *Block{
	var (
		virtualBlock *Block
	)

	virtualBlock = &Block{
		Index: -1,
		Hash: "",
	}
	return CreateNewBlock(virtualBlock, "Genesis Block")
}
