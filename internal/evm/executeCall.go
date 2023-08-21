package evm

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/paulrouge/xcall-event-watcher/internal/config"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common"
)

// Handles the reqIdAndData channel.
//
// Calls the 'executeCall' method on the xCall contract.
func CallExecuteCall() {
	for {
		r := <-ReqIdAndDataChan

		fmt.Println("reqId:", r.ReqId)
		// privateKeyHex := "YOUR_PRIVATE_KEY_HEX_WITHOUT_PREFIX"
		// var err error
		privateKey, err := crypto.HexToECDSA(privateKeyHex)
		if err != nil {
			log.Fatal(err)
		}

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
			fmt.Println(err)
		}

		toAddress := common.HexToAddress(config.SEPOLIA_XCALL_ADDRESS)

		// Encode the function call data using the ABI
		data, err := contractAbi.Pack("executeCall", r.ReqId, r.Data)
		if err != nil {
			log.Fatal(err)
		}

		tx := types.NewTransaction(nonce, toAddress, big.NewInt(0), gasLimit, gasPrice, data)
		chainID, err := EVMClient.NetworkID(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	
		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
		if err != nil {
			log.Fatal(err)
		}
	
		// err = EVMClient.SendTransaction(context.Background(), signedTx)
		// if err != nil {
		// 	log.Fatal(err)
		// }
	
		fmt.Printf("Transaction hash: %s\n", signedTx.Hash().Hex())

	}
}
