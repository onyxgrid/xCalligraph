# xCalligraph

<p align="center">

  <!-- <a href="https://godoc.org/github.com/onyxgrid/xCalligraph">
    <img src="https://godoc.org/github.com/onyxgrid/xCalligraph?status.svg" alt="GoDoc">
  </a> -->

  <a href="./LICENSE">
    <img src="https://img.shields.io/badge/License-MIT-blue.svg" alt="License: MIT">
  </a>

  <!-- make one for go ref stuff -->
  <a href="https://goreportcard.com/report/github.com/onyxgrid/xCalligraph">
    <img src="https://goreportcard.com/badge/github.com/onyxgrid/xCalligraph" alt="Go Report Card">
  </a>
</p>



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

## What could go wrong when using this
Keep in mind that you have to udpate the contract addresses in the .env file. Especially during testing this is something that is easily forgotten. If you don't update the contract addresses xCalligraph will not work.

## Issues
- [] for some reason call-messages are ignored when a bunch in a row are sent? In one example there were 3 callmessages in 1 block (block 4149075), only one (the first) got exectued by xCalligraph.

## TODO
- [] add rpc endpoint to .env file, should be possible to switch to any evm or jvm chain then?