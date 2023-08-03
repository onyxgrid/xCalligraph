package icon

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/icon-project/goloop/server/jsonrpc"
	v3 "github.com/icon-project/goloop/server/v3"
	"github.com/paulrouge/xcall-event-watcher/internal/logger"
	// TWO BELOW ARE FOR TESTING
	// "strconv"
	// "github.com/icon-project/goloop/server/jsonrpc"
)

type BlockHeightParam struct {
	Height string `json:"height" validate:"required,t_string"`
}

/*
Checks every block of the ICON blockchain and sends it to a channel.
The channel is used by handleBlock() to analyse each block and take out all transactions.
*/
func CheckBlocks() {

	latestBlock, err := Client.GetLastBlock()
	if err != nil {
		fmt.Println(err)
	}

	// FOR TESTING. STARTS HERE

	// _ = latestBlock
	// num := 11_663_037
	// hexString := strconv.FormatInt(int64(num), 16)

	// // num to jsonrpc.Hexint
	// numHex := jsonrpc.HexInt(hexString)

	// _ = numHex
	// blockHeightParam := &v3.BlockHeightParam{
	// 	Height: "0x" + numHex,
	// }

	// testBlock, err := Client.GetBlockByHeight(blockHeightParam)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// CurBlockChan <- testBlock
	// return

	// TESTING ENDS HERE

	for {
		currentBlock, _ := Client.GetLastBlock()
		
		// sleep 200 ms - to prevent blocking the port of the node?
		time.Sleep(300 * time.Millisecond)

		// check if there is a new block
		if currentBlock.Height > latestBlock.Height {
			
			// get the block 3 blocks before the current block, to prevent missing transactions/events
			heightHexString := strconv.FormatInt(int64(currentBlock.Height - 3), 16)
			heightHex := jsonrpc.HexInt(heightHexString)
			blockHeightParam := &v3.BlockHeightParam{
				Height: "0x" + heightHex,
			}
	
			blockToHandle, err := Client.GetBlockByHeight(blockHeightParam)
			
			if err != nil {
				fmt.Println(err)
				msg := fmt.Sprintf("GetBlockByHeight error: %v", err)
				logger.LogMessage(msg)
			}
	
			// fmt.Printf("Handling Block: %v\n", blockToHandle.Height)
	
			CurBlockChan <- blockToHandle
			latestBlock = currentBlock
		}
	}
}

// Handles the blocks sent by the CheckBlocks() function. Checks all transactions in the block.
func HandleBlock() {
	for {
		b := <-CurBlockChan

		// for testing. log the block height.
		// logger.Logger.Printf("Block %d", b.Height)

		for _, rawTx := range b.NormalTransactions {
			// rawTx is a []byte, convert to TransactionHashParam
			jsonTx := json.RawMessage(rawTx)
			tx := &v3.TransactionHashParam{}
			err := json.Unmarshal(jsonTx, tx)
			if err != nil {
				fmt.Println(err)
			}
			TransactionChan <- tx
		}
	}
}

/*
Gets the events of a transaction.
*/
func HandleTransaction() {
	for {
		tx := <-TransactionChan
		GetEvents(tx)
	}
}
