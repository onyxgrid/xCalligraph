package icon

import (
	"github.com/icon-project/goloop/client"
	v3 "github.com/icon-project/goloop/server/v3"
)

type reqIdAndData struct {
	ReqId string
	Data  string
}

var CurBlockChan = make(chan *client.Block, 1)
var TransactionChan = make(chan *v3.TransactionHashParam, 10)
var ReqIdAndDataChan = make(chan reqIdAndData, 1)
