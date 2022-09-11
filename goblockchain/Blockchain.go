package blockchain

import (
	"fmt"
	"container/list"
	"time"
	"crypto/sha256"
	"encoding/json"
)
type Block struct {
	Index int
	Timestamp string
	Transactions list.List
	PreviousHash string
	Nonce int
}

func (b Block) CalculateHash() string {
	blockData, _ := json.Marshal(b)
	hashInBytes := sha256.Sum256([]byte(blockData))
	return string(hashInBytes[:])
}

func (b *Block) BlockInit(index int, timestamp string, transactions list.List, previousHash string) Block {
	block := Block{index, timestamp, transactions, previousHash, 0}
	return block
}

type Blockchain struct {
	Chain list.List 
	GenesisBlock Block
	Difficulty int
	UnconfirmedTransactions list.List
}

func (b *Blockchain) CreateGenesisBlock() Block {
	genesisBlock := new(Block).BlockInit(0, time.Now().String(), list.List{}, "0")
	b.Chain.PushBack(genesisBlock)
	b.GenesisBlock = genesisBlock
	b.Difficulty = 2 // Difficulty of mining
	return genesisBlock
}

func (b *Blockchain) GetLatestBlockPointer() *Block {
	return b.Chain.Back().Value.(*Block)
}

func (b *Blockchain) GetLatestBlock() Block{
	return b.Chain.Back().Value.(Block)
}

func PrintChain(b *Blockchain){
	for e := b.Chain.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func (b *Blockchain) ProofOfWork(block Block) string{
	block.Nonce = 0
	hash := block.CalculateHash()

	for hash[:b.Difficulty] != FindPOFString(b.Difficulty) {
		block.Nonce++
		hash = block.CalculateHash()
		fmt.Println(hash)
	}
	return hash
}

func (b *Blockchain) AddBlock(block Block, hash string) bool {
	previousHash := b.GetLatestBlock().CalculateHash()
	if previousHash != block.PreviousHash {
		return false
	}
	if ! IsValidProof(b, block, hash) {
		return false
	}
	b.Chain.PushBack(&block)
	return true
}
//TODO: need to be fixed
type Transaction struct {
	sender string
	content string
	timestamp time.Time
}

func (b *Blockchain) AddNewTransaction(transaction Transaction) {
	b.UnconfirmedTransactions.PushBack(transaction)
}

func (b *Blockchain) Mine() int{
	if b.UnconfirmedTransactions.Len() == 0 {
		return 0
	}
	lastBlock := b.GetLatestBlock()
	newBlock := new(Block).BlockInit(lastBlock.Index + 1, time.Now().String(), b.UnconfirmedTransactions, lastBlock.CalculateHash())
	hash := b.ProofOfWork(newBlock)
	b.AddBlock(newBlock, hash)
	b.UnconfirmedTransactions.Init()
	return newBlock.Index
}
