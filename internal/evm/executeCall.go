package evm

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/paulrouge/xCalligraph/internal/config"
	"github.com/paulrouge/xCalligraph/internal/logger"
)

// Handles the reqIdAndData channel.
//
// Calls the 'executeCall' method on the xCall contract.
func CallExecuteCall() {
	
	for r := range ReqIdAndDataChan {
		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("error casting public key to ECDSA")
		}

		fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
		nonce, err := EVMClient.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			log.Fatal(err)
		}

		gasLimit := uint64(1000000)
		gasPrice, err := EVMClient.SuggestGasPrice(context.Background())
		if err != nil {
			fmt.Println("error at gasprice suggestion - ", err)
		}

		// add 10% to the suggested gas price
		gasPrice = gasPrice.Mul(gasPrice, big.NewInt(110))
		gasPrice = gasPrice.Div(gasPrice, big.NewInt(100))

		toAddress := common.HexToAddress(config.SEPOLIA_XCALL_ADDRESS)

		// fmt.Println("data - ", []byte(r.Data))
		// fmt.Println("reqId - ", r.ReqId)

		// Encode the function call data using the ABI
		data, err := contractAbi.Pack("executeCall", r.ReqId, r.Data)
		if err != nil {
			log.Fatal("error at abi encoding - ", err)
		}

		tx := types.NewTransaction(nonce, toAddress, big.NewInt(0), gasLimit, gasPrice, data)
		chainID, err := EVMClient.NetworkID(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
		if err != nil {
			log.Fatal("error at signing tx", err)
		}

		// if in test mode, just print the tx
		if config.TestMode {
			fmt.Printf("\nexecuteCall called on Sepolia xCall contract.\nreqId: 0x%s\ndata: %s\n", r.ReqId, r.Data)
			// sleep to prevent the api from crying
			time.Sleep(5 * time.Second)
			continue
		}

		err = EVMClient.SendTransaction(context.Background(), signedTx)
		if err != nil {
			log.Fatal(err)
		}

		// msg := "\nexecuteCall called on Sepolia xCall contract.\nreqId: 0x" + r.ReqId + "\ndata: " + r.Data + "\ntx: " + signedTx.Hash().Hex() + "\n"
		msg := fmt.Sprintf("\nexecuteCall called on Sepolia xCall contract.\nreqId: 0x%s\ndata: %s\ntx: %s\n", r.ReqId, r.Data, signedTx.Hash().Hex())
		logger.LogMessage(msg)
	}
}
