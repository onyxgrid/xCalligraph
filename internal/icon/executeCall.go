package icon

import (
	"fmt"
	"math/big"
	"github.com/eyeonicon/go-icon-sdk/transactions"
	"github.com/paulrouge/xcall-event-watcher/internal/logger"
)

// Send a transaction to the 'executeCall' method on the xCall contract.
func callExecuteCall(_reqId string, _data string) error {
	// the address of the contract
	method := "executeCall"

	// the params for the method,
	params := map[string]interface{}{
		"_reqId": _reqId,
		"_data":  _data,
	}

	value := big.NewInt(0)

	// We need to sign the tx, so we use the TransactionBuilder.
	tx := transactions.TransactionBuilder(Wallet.Address(), XCALL_ADDRESS, method, params, value)

	// sign the tx
	hash, err := Client.SendTransaction(Wallet, tx)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(*hash) // Returns the hash of the tx.

	logger.Logger.Printf("\nexecuteCall called on Berlin xCall contract.\nreqId: %v\ndata: %v\ntx: %v\n", _reqId, _data, *hash)

	return nil
}
