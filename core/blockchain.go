package core

import (
	"fmt"
)

type Blockchain struct {
	Blocks []*Block
}

func (bc *Blockchain) AppendBlock(newBlock *Block) {
	if len(bc.Blocks) == 0 {
		bc.Blocks = append(bc.Blocks, newBlock)
		return
	}
	if isValid(newBlock, bc.Blocks[len(bc.Blocks) - 1]) {
		bc.Blocks = append(bc.Blocks, newBlock)
	} else {
		panic("invalid block!!")
	}
}

func InitBlockChain() *Blockchain {
	genesisBlock := CreateGenesisBlock()
	blockChain := &Blockchain{}
	blockChain.AppendBlock(genesisBlock)
	return blockChain
}

func (bc *Blockchain) AppendData(data string) {
	preBlock := &bc.Blocks[len(bc.Blocks) - 1]
	newBlock := CreateNewBlock(*preBlock, data)
	bc.AppendBlock(newBlock)
}

func (bc *Blockchain) Print() {
	for _, block := range bc.Blocks {
		fmt.Printf("Index: %d\nHash:     %s\nPrevHash: %s\nData: %s\n\n", block.Index, block.Hash, block.PrevBlockHash, block.Data)
	}
}
func isValid(newBlock *Block, oldBlock *Block) bool {
	if newBlock.Index != oldBlock.Index + 1 {
		return false
	}

	if newBlock.PrevBlockHash != oldBlock.Hash {
		return false
	}

	if GetBlockHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}
