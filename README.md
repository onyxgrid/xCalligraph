# xCall event watcher
Tracks the ICON Blockchain for xCall events. Especially sendMessage events. Triggers a executeCall tx when a sendMessage event is found.

If you hook up a wallet with funds to the watcher it will automatically send the executeCall tx.

Atm it will be connected to the Berlin testnet. So it will only work with Berlin testnet as destination chain.

## How to use
1. Change the .env.example to .env and fill in the required values.
2. Place a wallet keystore file in the wallet folder and name it `keystore`.

## How to run
Make sure to have Docker and docker-compose installed. Then simply run `make run` in the root folder.

