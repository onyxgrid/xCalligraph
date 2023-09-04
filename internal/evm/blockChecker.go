package evm

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/paulrouge/xCalligraph/internal/config"
)

func CheckBlocks() {
	latestBlock, _ := EVMClient.BlockNumber(context.Background())

	if config.TestMode {
		_ = latestBlock
		num := config.SepoliaTestBlockHeight
		bigNum := big.NewInt(int64(num))

		testBlock, err := EVMClient.BlockByNumber(context.Background(), bigNum)
		if err != nil {
			fmt.Println(err)
		}

		CurBlockChan <- testBlock
		return
	}

	for {
		currentBlock, _ := EVMClient.BlockNumber(context.Background())

		// sleep 200 ms - to prevent blocking the port of the node?
		time.Sleep(3 * time.Second)

		// check if there is a new block
		if currentBlock > latestBlock {
			// get the block 3 blocks before the current block, to prevent missing transactions/events
			blockToHandle := currentBlock - 3

			blockBigInt := new(big.Int)
			blockBigInt.SetUint64(blockToHandle)
			block, err := EVMClient.BlockByNumber(context.Background(), blockBigInt)
			if err != nil {
				fmt.Println(err)
			}

			CurBlockChan <- block
			latestBlock = currentBlock
		}

	}

}

// Handles the blocks sent by the CheckBlocks() function. Checks all transactions in the block.
func HandleBlock() {
	for b := range CurBlockChan {
		// for testing. log the block height.
		fmt.Printf("Block %d\n", b.NumberU64())
		fmt.Printf("Amount of transactions: %d\n", len(b.Transactions()))

		for _, tx := range b.Transactions() {
			// to := tx.To()
			// fmt.Println("to:", to.Hex())
			// fmt.Println(tx)
			TransactionChan <- tx
		}
	}
}

// Handles the transactions sent by the HandleBlock() function. Checks if the transaction is a xCall transaction.
func HandleTransaction() {
	for tx := range TransactionChan {
		if tx.To() == nil {
			continue
		}

		if tx.To().Hex() == config.SEPOLIA_BMC_ADDRESS {
			EVMGetEvents(tx.Hash().Hex())
		}
	}
}
