package icon

import (
	"fmt"

	"github.com/icon-project/goloop/client"
	v3 "github.com/icon-project/goloop/server/v3"
)

const (
	BTP_EVENT         = "BTPEvent(str,int,str,str)"
	CALLMESSAGE_EVENT = "CallMessage(str,str,int,int,bytes)"
)

func GetEvents(_hash *v3.TransactionHashParam) error {
	tx, err := Client.GetTransactionResult(_hash)
	if err != nil {
		return fmt.Errorf("GetTransactionByHash error: %v", err)
	}

	// height, _ := tx.BlockHeight.Int64()
	// fmt.Println("checking events of txs in block:", height)

	logs := tx.EventLogs

	for _, log := range logs {
		address := log.Addr

		if address.Address().String() == BMC_ADDRESS {
			HandleBmcEvents(log)
		}
		if address.Address().String() == XCALL_ADDRESS {
			HandleXCallEvents(log)
		}
	}

	return nil
}

// Handle the eventlogs send by the xcall contract
func HandleXCallEvents(_log client.EventLog) {
	if *_log.Indexed[0] == CALLMESSAGE_EVENT {

		xCallRequestID := *_log.Data[0]
		xCallData := *_log.Data[1]
		btpAddressCaller := *_log.Indexed[1]

		if btpAddressCaller == BTP_ADDRESS_TO_TRACK {
			err := callExecuteCall(xCallRequestID, xCallData)
			if err != nil {
				fmt.Println("callExecuteCall error:", err)
			}
		}
	}
}

// Handle the eventlogs send by the bmc contract
func HandleBmcEvents(_log client.EventLog) {
	// if *_log.Indexed[0] == BTP_EVENT {
	// 	fmt.Println("BTP_EVENT")
	// }
}
