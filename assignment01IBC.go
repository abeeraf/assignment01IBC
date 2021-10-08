package assignment01IBC

import (
	"crypto/sha256"
	"fmt"
)

type BlockData struct {
	Transactions []string
}

type Block struct {
	Data        BlockData
	PrevPointer *Block
	PrevHash    string
	CurrentHash string
}

func CalculateHash(inputBlock *Block) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%v", inputBlock.Data.Transactions))))
}

func InsertBlock(dataToInsert BlockData, chainHead *Block) *Block {

	var newBlock *Block = new(Block)
	newBlock.Data = dataToInsert
	if chainHead != nil {
		newBlock.PrevHash = chainHead.CurrentHash
		newBlock.PrevPointer = chainHead
	} else {
		newBlock.PrevHash = ("0")
		newBlock.PrevPointer = nil
	}
	newBlock.CurrentHash = CalculateHash(newBlock)

	return newBlock
}

func ChangeBlock(oldTrans string, newTrans string, chainHead *Block) {

	for trav := chainHead; trav != nil; trav = trav.PrevPointer {
		for i := 0; i < len(trav.Data.Transactions); i++ {
			if trav.Data.Transactions[i] == oldTrans {
				trav.Data.Transactions[i] = newTrans
				trav.CurrentHash = CalculateHash(trav)
			}
		}
	}
}

func ListBlocks(chainHead *Block) {

	for trav := chainHead; trav != nil; trav = trav.PrevPointer {
		fmt.Println("\n------Block------")
		fmt.Println("Data:        ", trav.Data)
		fmt.Println("PrevHash:    ", trav.PrevHash)
		fmt.Println("CurrentHash: ", trav.CurrentHash)
		fmt.Println()
	}

}

func VerifyChain(chainHead *Block) {

	prev := chainHead.PrevPointer

	for trav := chainHead; trav.PrevPointer != nil; trav = trav.PrevPointer {
		if prev.CurrentHash != trav.PrevHash {
			fmt.Println("Chain Compromised!")
			break
		}
		prev = chainHead.PrevPointer
	}

}
