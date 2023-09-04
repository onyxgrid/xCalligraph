
<p align="center">
  <a href="./LICENSE">
    <img src="https://img.shields.io/badge/License-MIT-blue.svg" alt="License: MIT">
  </a>

  <!-- make one for go ref stuff -->
  <a href="https://goreportcard.com/report/github.com/onyxgrid/xCalligraph">
    <img src="https://goreportcard.com/badge/github.com/onyxgrid/xCalligraph" alt="Go Report Card">
  </a>
</p>

# xCalligraph

xCalligraph tracks the ICON Berlin Testnet, and the Ethereum Sepolia Testnet for xCall events. Especially sendMessage events, and triggers a executeCall transaction when a sendMessage event is found.

If you hook up a wallet with funds to xCalligraphm it will automatically send the executeCall transaction.

## Before running
1. Change the .env.example to .env and fill in the required values.
2. Place a wallet keystore file in the wallet folder and name it `keystore` (Without file extension like .json .) Make sure that the wallet has funds on the Berlin testnet. If want to use xCalligraph to automatically send the executeCall transaction on Sepolia, make sure to have funds on the Sepolia testnet as well and put the private key of that wallet in the `.env` file.

So make sure that you have set up the `.env` file correctly for the chain you need to sign the xCall `executeCall` function on. 

### the .env file
```
BERLIN_WALLET_PASSWORD=password to the keystorefile at wallet/keystore (without file extension)
SEPOLIA_PRIVATE_KEY=private key of the wallet that will sign the executeCall transaction on Sepolia

# addresses (dApps) to monitor, comma separated, all chains. keep on one line.
BTP_ADDRESS_TO_TRACK=btp://0xaa36a7.eth2/0x5c4fA4b22256Ff15E5A1aa02517d07d17cF7A7bE,btp://0x7.icon/cx3723d8cb8d8ac7da29f692ce2abc8156423631be

BERLIN_XCALL_ADDRESS=cxf4958b242a264fc11d7d8d95f79035e35b21c1bb
BERLIN_BMC_ADDRESS=cxf1b0808f09138fffdb890772315aeabb37072a8a
SEPOLIA_BMC_ADDRESS=0xE602326106f5E1d436a3CCEB2A408759925f81ff
SEPOLIA_XCALL_ADDRESS=0x694C1f5Fb4b81e730428490a1cE3dE6e32428637

# if test mode is true make sure to set both test blockheigts to a valid int64 
# test blockheights must be numbers without commas, dots, or underscores. 
# So 1000000 is valid, 1_000_000 is not, nor is 1.000.000 etc.
BERLIN_TEST_BLOCKHEIGHT=1000000
SEPOLIA_TEST_BLOCKHEIGHT=1000000

TEST_WITH_SIGNING_TXS=false # set to true if you want to send the executeCall transactions
```

_BTP_ADDRESS_TO_TRACK_ holds the BTP-addresses of the contract. These are the addresses that will be tracked for events. And only events from this address will call the `executeCall` xCall function.

## How to run
Make sure to have Docker and docker-compose installed. Then simply run `make run` in the root folder. This will build the docker image and start the container in the background of your system. 

To stop the container run `make stop`.

# How to run in test mode
In testmode you can run it with or without transaction signing. If you run it without signing it will only print the tx data to the console. If you run it with signing it will send the signed 'executeCall' transaction. The transaction will then also be logged to the transactions.log file.

Before you run a test, make sure to set the blocknumber you want to check for the xCall event in the `.env` file.

To run in test mode without signing run `make test`. To run in test mode with signing run `make testwithsigning`. Needless to say that you need to have the correct wallet config set up in the `.env` file as well.

## What could go wrong when using this
Keep in mind that you have to update the contract addresses in the `.env` file. Especially during testing this is something that is easily forgotten. If you don't update the contract addresses xCalligraph will not work.

## How to contribute
If you want to contribute to this project, feel free to fork it and create a pull request. If you have any questions, feel free to open an issue.