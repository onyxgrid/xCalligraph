package icon

import (
	"fmt"
	"math/big"

	"github.com/eyeonicon/go-icon-sdk/transactions"
	"github.com/paulrouge/xCalligraph/internal/config"
	"github.com/paulrouge/xCalligraph/internal/logger"
)

// i don't think this is needed actually
var handledReqIDs []string

// Handles the reqIdAndData channel.
//
// Calls the 'executeCall' method on the xCall contract.
func CallExecuteCall() {
	for r := range ReqIdAndDataChan {
		for _, reqID := range handledReqIDs {
			if reqID == r.ReqId {
				return
			}
		}

		// the method of the contract
		method := "executeCall"

		// the params for the method,
		params := map[string]interface{}{
			"_reqId": r.ReqId,
			"_data":  r.Data,
		}

		value := big.NewInt(0)

		// We need to sign the tx, so we use the TransactionBuilder.
		tx := transactions.TransactionBuilder(Wallet.Address(), config.XCALL_ADDRESS, method, params, value)

		handledReqIDs = append(handledReqIDs, r.ReqId)

		// if we are in test mode, just print the tx
		if config.TestMode {
			fmt.Printf("\nexecuteCall called on Berlin xCall contract.\nreqId: %v\ndata: %v\n", r.ReqId, r.Data)
			continue
		}

		// sign the tx
		hash, err := Client.SendTransaction(Wallet, tx)
		if err != nil {
			fmt.Println(err)
		}

		// fmt.Println(*hash) // Returns the hash of the tx.
		msg := fmt.Sprintf("\nexecuteCall called on Berlin xCall contract.\nreqId: %v\ndata: %v\ntx: %v\n", r.ReqId, r.Data, *hash)
		logger.LogMessage(msg)
	}
}
