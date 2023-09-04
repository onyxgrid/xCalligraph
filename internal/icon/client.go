package icon

import (
	"fmt"
	"log"
	"os"

	"github.com/eyeonicon/go-icon-sdk/networks"
	"github.com/eyeonicon/go-icon-sdk/wallet"
	"github.com/icon-project/goloop/client"
	"github.com/icon-project/goloop/module"
	"github.com/joho/godotenv"
)

var (
	Client *client.ClientV3 = nil
	Wallet module.Wallet    = nil
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Client = GetClient("berlin")

	PASSWORD := os.Getenv("BERLIN_WALLET_PASSWORD")

	if PASSWORD == "" {
		panic("WALLET_PASSWORD is empty. Please check your .env file.")
	}

	// err as a seperate var so we do not have to redeclare Wallet by using :=
	var _err error
	Wallet, _err = wallet.LoadWallet("./wallet/keystore", PASSWORD)
	if _err != nil {
		panic(_err)
	}
}

// Call with 'Berlin' or empty for main.
func GetClient(args ...string) *client.ClientV3 {

	if len(args) == 0 {
		fmt.Println("Connecting to mainnet via EOI Node endpoint...")
		networks.SetActiveNetwork(networks.Mainnet())
		// Client := client.NewClientV3(networks.GetActiveNetwork().URL)
		Client := client.NewClientV3("http://65.21.202.45:9000/api/v3")
		return Client
	}

	if args[0] == "berlin" {
		// fmt.Println("Connecting to Berlin...")

		berlinNetwork := networks.Network{
			URL: "https://berlin.net.solidwallet.io/api/v3",
			NID: "0x7",
		}

		networks.SetActiveNetwork(berlinNetwork)
		Client := client.NewClientV3(networks.GetActiveNetwork().URL)
		return Client
	} else {
		networks.SetActiveNetwork(networks.Mainnet())
		Client := client.NewClientV3(networks.GetActiveNetwork().URL)
		return Client
	}
}
