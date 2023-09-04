package evm

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

type reqIdAndData struct {
	ReqId *big.Int
	Data  []byte
}

var CurBlockChan = make(chan *types.Block, 1)
var TransactionChan = make(chan *types.Transaction)
var ReqIdAndDataChan = make(chan reqIdAndData)
