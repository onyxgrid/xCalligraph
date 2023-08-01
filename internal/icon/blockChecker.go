package icon

import (
	"encoding/json"
	"fmt"
	"time"
	v3 "github.com/icon-project/goloop/server/v3"
	
	// TWO BELOW ARE FOR TESTING
	"strconv"
	"github.com/icon-project/goloop/server/jsonrpc"
)

type BlockHeightParam struct {
	Height string `json:"height" validate:"required,t_string"`
}

/*
Checks every block of the ICON blockchain and sends it to a channel.
The channel is used by Handle to analyse each block and take out all transactions.
*/
func CheckBlocks() {

	latestBlock, err := Client.GetLastBlock()
	if err != nil {
		fmt.Println(err)
	}

	// FOR TESTING. STARTS HERE

	_ = latestBlock
	num := 11_493_046
	hexString := strconv.FormatInt(int64(num), 16)

	// // num to jsonrpc.Hexint
	numHex := jsonrpc.HexInt(hexString)

	_ = numHex
	blockHeightParam := &v3.BlockHeightParam{
		Height: "0x" + numHex,
	}

	testBlock, err := Client.GetBlockByHeight(blockHeightParam)
	if err != nil {
		fmt.Println(err)
	}

	CurBlockChan <- testBlock
	return
	
	// TESTING ENDS HERE
	
	for {
		currentBlock, _ := Client.GetLastBlock()

		// sleep 200 ms - to prevent blocking the port of the node?
		time.Sleep(300 * time.Millisecond)

		// check if there is a new block
		if currentBlock.Height > latestBlock.Height {

			CurBlockChan <- currentBlock
			// set latestBlock to currentBlock
			latestBlock = currentBlock
		}
	}

}

/*
Handles the blocks sent by the CheckBlocks() function. Checks all transactions in the block.
*/
func HandleBlock() {
	for {
		b := <-CurBlockChan
		fmt.Printf("Block %d\n", b.Height)
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
