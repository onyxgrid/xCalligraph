
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

xCalligraph tracks the ICON Blockchain for xCall events. Especially sendMessage events. Triggers a executeCall tx when a sendMessage event is found.

If you hook up a wallet with funds to xCalligraph it will automatically send the executeCall tx.

Atm it will be connected to the Berlin testnet. So it will only work with Berlin testnet as destination chain.

## Before running
1. Change the .env.example to .env and fill in the required values.
2. Place a wallet keystore file in the wallet folder and name it `keystore` (Without file extension like .json .) Make sure that the wallet has funds on the Berlin testnet.

### the .env file
```
WALLET_PASSWORD=passwordhere
BTP_ADDRESS_TO_TRACK=btp://a_valid/btp_address
BERLIN_XCALL_ADDRESS=cxf4958b242a264fc11d7d8d95f79035e35b21c1bb
BERLIN_BMC_ADDRESS=cxf1b0808f09138fffdb890772315aeabb37072a8a
```

_BTP_ADDRESS_TO_TRACK_ is the BTP address of the xCall contract on the source chain. This is the address that will be tracked for events. And only events from this address will call the `executeCall` function on the Berlin chain.

## How to run
Make sure to have Docker and docker-compose installed. Then simply run `make run` in the root folder. This will build the docker image and start the container in the background of your system. 

To stop the container run `make stop`.

# How to run in test mode
In testmoded you can run it with or without transaction signing. If you run it without signing it will only print the tx data to the console. If you run it with signing it will send the signed 'executeCall' transaction. The transaction will then also be logged to the transactions.log file.

Before you run a test, make sure to set the blocknumber you want to check for the xCall event in the `.env` file.

To run in test mode without signing run `make test`. To run in test mode with signing run `make test-sign`. Needless to say that you need to have the correct wallet config set up in the `.env` file as well.

## What could go wrong when using this
Keep in mind that you have to update the contract addresses in the `.env` file. Especially during testing this is something that is easily forgotten. If you don't update the contract addresses xCalligraph will not work.

## How to contribute
If you want to contribute to this project, feel free to fork it and create a pull request. If you have any questions, feel free to open an issue.