package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type BlockChain struct {
	blocks []*Block
}

type Block struct {
	hash []byte
	data []byte
	prev []byte
}

func (b *Block) Hash() {
	info := bytes.Join([][]byte{b.data, b.prev}, []byte{})
	hash := sha256.Sum256(info)
	b.hash = hash[:]
}

func CreateBlock(data string, prev []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prev}
	block.Hash()
	return block
}

func (ch *BlockChain) AddBlock(data string) {
	prev := ch.blocks[len(ch.blocks)-1]
	new := CreateBlock(data, prev.hash)
	ch.blocks = append(ch.blocks, new)
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{CreateBlock("Genesis", []byte{})}}
}

func main() {
	chain := InitBlockChain()

	chain.AddBlock("block 1")
	chain.AddBlock("block 2")
	chain.AddBlock("block 3")
	chain.AddBlock("block 4")

	for _, b := range chain.blocks {
		fmt.Printf("prev hash: %x\n", b.prev)
		fmt.Printf("block data: %s\n", b.data)
		fmt.Printf("hash: %x\n", b.hash)
	}
}
