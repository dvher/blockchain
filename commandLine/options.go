package commandLine

import (
	"blockchain1/blockchain"
	"fmt"
	"strconv"
)

func createBlockChain(address string){
	chains := blockchain.InitBlockChain(address)
	defer chains.Database.Close()
	fmt.Println("Blockchain Created")
}

func getBalance(address string) {
	chain := blockchain.ContinueBlockChain(address)
	defer chain.Database.Close()

	balance := 0
	UTXOs := chain.FindUTXO(address)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of %s: %d\n", address, balance)
}

func printChain() {
	chains := blockchain.ContinueBlockChain("")
	defer chains.Database.Close()
	iter := chains.Iterator()

	for {
		block := iter.Next()

		fmt.Printf("Previous hash: %x\n", block.PreviousHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PreviousHash) == 0 {
			break
		}
	}
}

func send(from, to string, amount int) {
	chains := blockchain.ContinueBlockChain(from)
	defer chains.Database.Close()

	tx := blockchain.NewTransaction(from, to, amount, chains)
	chains.AddBlock([]*blockchain.Transaction{tx})
	fmt.Println("Success sending coins")
}

func searchBlockByHash(blockHash string) {
	chain := blockchain.ContinueBlockChain("")
	defer chain.Database.Close()

	iter := chain.Iterator()
	for {
		block := iter.Next()
		if block == nil {
			fmt.Println("Bloque no encontrado.")
			break
		}

		if fmt.Sprintf("%x", block.Hash) == blockHash {
			fmt.Println("Bloque encontrado!")
			break
		}
	}
}


