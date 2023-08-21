package evm

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/paulrouge/xcall-event-watcher/internal/config"
	"golang.org/x/crypto/sha3"
)

// var Hashwithxcallevent = "0x227ac4cc33c877735370e5824bbdf2a9b3ede4d4208c30dc7951a6977f4e78cf"
// var HashCallMessageSent = "0x5a6d720ad4718ad86b624dea2fc0a2edde44cebabcc14dcc7a08151c99171662" // not executed yet
// var Hashexecutecall = "0xeaacea77d9469c0e775adce5974be7da775f24f70bd36e7d3d651f56611b9210"
// var Hashwithotherlog = "0x3d6a3164ea0ccbeb901d33d74d3e65c5338e0b6e59f33545b6e7295c7ade7b74"
// var BTPADDRESSBERLINDAPP = "btp://0x7.icon/cx39bf06738279054733580d179ce6eab0ed19a8c2"

func EVMGetEvents(_hash string) {

	ctx := context.Background()
	s := common.HexToHash(_hash)

	tx, err := EVMClient.TransactionReceipt(ctx, s)
	if err != nil {
		panic(err)
	}

	logs := tx.Logs

	for _, log := range logs {

		_fromBTPAddress := log.Topics[1]

		for _, btpAddr := range config.BTP_ADDRESSES_TO_TRACK {
			sha3Hash := sha3.NewLegacyKeccak256()
			sha3Hash.Write([]byte(btpAddr))
			fromBTPAddressSHA3 := common.BytesToHash(sha3Hash.Sum(nil))

			if _fromBTPAddress == fromBTPAddressSHA3 {
				t, err := contractAbi.Unpack("CallMessage", log.Data)
				_ = t
				// if err != nil, it's not a xCall event, return
				if err != nil {
					return
				}
				_reqId := t[0]
				_data := t[1].([]byte) // this is a []uint8/[]byte type

				// fmt.Println("dataIneed:", string(_data))

				newReqIdAndData := reqIdAndData{
					ReqId: "0x" + _reqId.(*big.Int).Text(16),
					Data:  string(_data),
				}
				// fmt.Println("got here")
				// _ = newReqIdAndData

				ReqIdAndDataChan <- newReqIdAndData

			}
		}
	}

}
