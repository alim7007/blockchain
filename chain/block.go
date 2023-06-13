package chain

import (
	"bytes"
	"crypto/sha256"
)

type Block struct {
	Hash     []byte
	PrevHash []byte
	Data     []byte
}

type BlockChain struct {
	Blocks []*Block
}

func (block *Block) DeriveHash() {
	// fmt.Printf("%q\n",[][]byte{block.Data,block.PrevHash})
	info := bytes.Join([][]byte{block.Data, block.PrevHash}, []byte{})
	// fmt.Printf("%q\n",info)
	hash := sha256.Sum256(info)
	// fmt.Printf("%q\n",hash)
	block.Hash = hash[:]
}

func createBlock(data string, prevHash []byte) *Block {
	// creating address of block
	block := &Block{[]byte{}, prevHash, []byte(data)}
	// sending address of block
	block.DeriveHash()
	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := createBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

func genesis() *Block {
	return createBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{genesis()}}
}
