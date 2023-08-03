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


// Get all events from the transaction.
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
		
		// logger.Logger.Printf("Event from: %v", address.Address().String())

		if address.Address().String() == BMC_ADDRESS {
			HandleBMCEvents(log)
		}
		if address.Address().String() == XCALL_ADDRESS {
			HandleXCallEvents(log, address.Address().String(), )
		}
	}

	return nil
}

// Handle the eventlogs send by the xcall contract.
// 
// If the event is a 'CallMessage(str,str,int,int,bytes)' event, 
// we call the 'executeCall' method on the xCall contract.
func HandleXCallEvents(_log client.EventLog, _testAdr string) {
	
	eventName := *_log.Indexed[0]

	if eventName == CALLMESSAGE_EVENT {

		xCallRequestID := *_log.Data[0]
		xCallData := *_log.Data[1]
		btpAddressCaller := *_log.Indexed[1]
		
		if btpAddressCaller == BTP_ADDRESS_TO_TRACK {
			newReqIdAndData := reqIdAndData{
				ReqId: xCallRequestID,
				Data: xCallData,
			}

			ReqIdAndDataChan <- newReqIdAndData
		}
	}
}

// Handle the eventlogs send by the bmc contract.
func HandleBMCEvents(_log client.EventLog) {
	// if *_log.Indexed[0] == BTP_EVENT {
	// 	fmt.Println("BTP_EVENT")
	// }
}
