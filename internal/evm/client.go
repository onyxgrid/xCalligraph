package evm

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

var EVMClient *Client

// Client is a wrapper around the Ethereum client.
type Client struct {
	*ethclient.Client
}

// NewClient creates a new Ethereum client.
func NewClient() (*Client, error) {
	client, err := ethclient.Dial("https://ethereum-sepolia-rpc.allthatnode.com")
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
} 

func init(){
	var err error
	EVMClient, err = NewClient()
	if err != nil {
		panic(err)
	}
}
