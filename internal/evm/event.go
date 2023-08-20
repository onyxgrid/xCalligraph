package evm

import (
	"context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/paulrouge/xcall-event-watcher/internal/icon"
	"golang.org/x/crypto/sha3"
)

var Hashwithxcallevent = "0x227ac4cc33c877735370e5824bbdf2a9b3ede4d4208c30dc7951a6977f4e78cf"
var HashCallMessageSent = "0x5a6d720ad4718ad86b624dea2fc0a2edde44cebabcc14dcc7a08151c99171662" // not executed yet
var Hashexecutecall = "0xeaacea77d9469c0e775adce5974be7da775f24f70bd36e7d3d651f56611b9210"
var Hashwithotherlog = "0x3d6a3164ea0ccbeb901d33d74d3e65c5338e0b6e59f33545b6e7295c7ade7b74"
var BTPADDRESSBERLINDAPP = "btp://0x7.icon/cx39bf06738279054733580d179ce6eab0ed19a8c2"

func EVMGetEvents() {
	ctx := context.Background()
	s := common.HexToHash(HashCallMessageSent)

	tx, err := EVMClient.TransactionReceipt(ctx, s)
	if err != nil {
		panic(err)
	}

	// xCall contract ABI
	var contractABI = `[{"anonymous":false,"inputs":[{"indexed":true,"internalType":"uint256","name":"_reqId","type":"uint256"},{"indexed":false,"internalType":"int256","name":"_code","type":"int256"},{"indexed":false,"internalType":"string","name":"_msg","type":"string"}],"name":"CallExecuted","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"string","name":"_from","type":"string"},{"indexed":true,"internalType":"string","name":"_to","type":"string"},{"indexed":true,"internalType":"uint256","name":"_sn","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"_reqId","type":"uint256"},{"indexed":false,"internalType":"bytes","name":"_data","type":"bytes"}],"name":"CallMessage","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"_from","type":"address"},{"indexed":true,"internalType":"string","name":"_to","type":"string"},{"indexed":true,"internalType":"uint256","name":"_sn","type":"uint256"},{"indexed":false,"internalType":"int256","name":"_nsn","type":"int256"}],"name":"CallMessageSent","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"uint8","name":"version","type":"uint8"}],"name":"Initialized","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"uint256","name":"_sn","type":"uint256"},{"indexed":false,"internalType":"int256","name":"_code","type":"int256"},{"indexed":false,"internalType":"string","name":"_msg","type":"string"}],"name":"ResponseMessage","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"uint256","name":"_sn","type":"uint256"},{"indexed":false,"internalType":"int256","name":"_code","type":"int256"},{"indexed":false,"internalType":"string","name":"_msg","type":"string"}],"name":"RollbackExecuted","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"uint256","name":"_sn","type":"uint256"}],"name":"RollbackMessage","type":"event"},{"inputs":[],"name":"admin","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"_reqId","type":"uint256"},{"internalType":"bytes","name":"_data","type":"bytes"}],"name":"executeCall","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"_sn","type":"uint256"}],"name":"executeRollback","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"getBtpAddress","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"string","name":"_net","type":"string"},{"internalType":"bool","name":"_rollback","type":"bool"}],"name":"getFee","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getProtocolFee","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getProtocolFeeHandler","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"string","name":"_src","type":"string"},{"internalType":"string","name":"_svc","type":"string"},{"internalType":"uint256","name":"_sn","type":"uint256"},{"internalType":"uint256","name":"_code","type":"uint256"},{"internalType":"string","name":"_msg","type":"string"}],"name":"handleBTPError","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"_from","type":"string"},{"internalType":"string","name":"_svc","type":"string"},{"internalType":"uint256","name":"_sn","type":"uint256"},{"internalType":"bytes","name":"_msg","type":"bytes"}],"name":"handleBTPMessage","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"_bmc","type":"address"}],"name":"initialize","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"_to","type":"string"},{"internalType":"bytes","name":"_data","type":"bytes"},{"internalType":"bytes","name":"_rollback","type":"bytes"}],"name":"sendCallMessage","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"payable","type":"function"},{"inputs":[{"internalType":"address","name":"_address","type":"address"}],"name":"setAdmin","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"_value","type":"uint256"}],"name":"setProtocolFee","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"_addr","type":"address"}],"name":"setProtocolFeeHandler","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"toAddr","type":"address"},{"internalType":"string","name":"to","type":"string"},{"internalType":"string","name":"from","type":"string"},{"internalType":"bytes","name":"data","type":"bytes"}],"name":"tryHandleCallMessage","outputs":[],"stateMutability":"nonpayable","type":"function"}]`
	contractAbi, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		panic(err)
	}

	logs := tx.Logs

	for _, log := range logs {

		_fromBTPAddress := log.Topics[1]

		for _, btpAddr := range icon.BTP_ADDRESSES_TO_TRACK {
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
				fmt.Println("registrered btp address")
				_reqId := t[0]
				_dataIneed := t[1].([]byte)
				fmt.Println("_reqId", _reqId)
				fmt.Println("_data:" , string(_dataIneed))
			}	
		}
	}

}
