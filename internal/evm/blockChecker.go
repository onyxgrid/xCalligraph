package evm

import (
	"context"
	"fmt"
	"math/big"
	"time"
)

func CheckBlocks() {

	latestBlock, _ := EVMClient.BlockNumber(context.Background())
	fmt.Println("currentBlock:", latestBlock)

	for {
		currentBlock, _ := EVMClient.BlockNumber(context.Background())

		// sleep 200 ms - to prevent blocking the port of the node?
		time.Sleep(300 * time.Millisecond)

		// check if there is a new block
		if currentBlock > latestBlock {
			// get the block 3 blocks before the current block, to prevent missing transactions/events
			blockToHandle := currentBlock - 3

			fmt.Printf("Handling Block: %v\n", blockToHandle)

			blockBigInt := new(big.Int)
			blockBigInt.SetUint64(blockToHandle)

			block, err := EVMClient.BlockByNumber(context.Background(), blockBigInt)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Printf("Handling Txs: %v\n", block.Transactions())

			// CurBlockChan <- blockToHandle
			latestBlock = currentBlock
		}

	}

}
