package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

// TODO => the block type consists of {hash, data, prevHash}
type Block struct {
	Data     []byte // Actual transactions data
	Hash     []byte // the hash is the hashing result of the entire block fields (Data + PrevHash)
	PrevHash []byte // the hash of the prev block
}

// HashBlock generates a hash of a receiver block by joining the Data and PrevHash and hash the result via sha256
func (b *Block) HashBlock() {
	// join the Data and PrevHash together
	data := bytes.Join([][]byte{
		b.Data,
		b.PrevHash,
	}, []byte{})
	// hash the data
	hash := sha256.Sum256(data)
	// set the Hash of the current block
	b.Hash = hash[:]
}

// create a new block given a data string and the previous block's hash
func NewBlock(data string, prevBlockHash []byte) *Block {
	b := &Block{Hash: []byte{}, Data: []byte(data), PrevHash: prevBlockHash}
	b.HashBlock()
	return b
}

// TODO => the blockchain type consists of multiple blocks
type Blockchain struct {
	blocks []*Block
}

// create a root block (known as gensis)
func GensisBlock() *Block {
	data := []byte{}
	prevHash := []byte{}
	gensis := NewBlock(string(data), prevHash)
	return gensis
}

// create new blockchain with gensis root block
func NewBlockchain() *Blockchain {
	gensis_root_block := GensisBlock()
	return &Blockchain{
		blocks: []*Block{gensis_root_block},
	}
}

// add a block to current chain
func (bc *Blockchain) AddBlockToChain(data string) {
	// get the previous block and fetch its Hash
	prev := bc.blocks[len(bc.blocks)-1]
	prevHash := prev.Hash

	// create a block with this data and the prev block hash
	newBlock := NewBlock(data, prevHash)

	// append the new block to the chain
	bc.blocks = append(bc.blocks, newBlock)
}

// to explore the blockchain
func (bc Blockchain) Explore() {
	for i, b := range bc.blocks {
		fmt.Printf("BLOCK NUMBER.%v\n", i)
		fmt.Printf("\t Block With Hash \t ➜ %v \n", b.Hash)
		fmt.Printf("\t Block With Data \t ➜ %v \n", b.Data)
		fmt.Printf("\t Block With PrevHash \t ➜ %v \n", b.PrevHash)
	}
}
func main() {
	// gensis := GensisBlock()
	// block_1 := NewBlock("sending tx_1 from f to m", gensis.Hash)
	// block_2 := NewBlock("sending tx_2 from m to f", block_1.Hash)
	// block_3 := NewBlock("sending tx_3 from m to f", block_2.Hash)
	// blockchain := NewBlockchain()
	blockchain := NewBlockchain()
	blockchain.AddBlockToChain("sending tx_1 from f to m")
	blockchain.AddBlockToChain("sending tx_2 from a to m")
	blockchain.AddBlockToChain("sending tx_3 from x to m")
	blockchain.Explore()
}
