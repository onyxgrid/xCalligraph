package icon

import (
	"github.com/icon-project/goloop/client"
	v3 "github.com/icon-project/goloop/server/v3"
)

var CurBlockChan = make(chan *client.Block, 3)
var TransactionChan = make(chan *v3.TransactionHashParam, 10)

