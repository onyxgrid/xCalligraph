package evm

import "github.com/ethereum/go-ethereum/core/types"

type reqIdAndData struct {
	ReqId string
	Data  string
}

var CurBlockChan = make(chan *types.Block, 1)
var TransactionChan = make(chan *types.Transaction, 20)
var ReqIdAndDataChan = make(chan reqIdAndData, 1)
