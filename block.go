package main

import (
	"time"
)
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

func NewBlock(data string, preBlockHash []byte) *Block{
		block := &Block{time.Now().Unix(), []byte(data), preBlockHash, []byte{},0}
		pow := NewProofOfWork(block)
		nonce, hash := pow.Run()

		block.Hash = hash[:]
		block.Nonce = nonce
		return block
}
