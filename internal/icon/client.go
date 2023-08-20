package icon

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/eyeonicon/go-icon-sdk/networks"
	"github.com/eyeonicon/go-icon-sdk/wallet"
	"github.com/icon-project/goloop/client"
	"github.com/icon-project/goloop/module"
	"github.com/joho/godotenv"
	"github.com/paulrouge/xcall-event-watcher/internal/config"
)

var (
	Client                 *client.ClientV3 = nil
	Wallet                 module.Wallet    = nil
	BTP_ADDRESSES_TO_TRACK []string
	XCALL_ADDRESS          string
	BMC_ADDRESS            string
	RELAY_ADDRESS          string
)

func init() {
	Client = GetClient("berlin")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PASSWORD := os.Getenv("WALLET_PASSWORD")
	BTP_ADDRESS_LINE := os.Getenv("BTP_ADDRESS_TO_TRACK")
	XCALL_ADDRESS = os.Getenv("BERLIN_XCALL_ADDRESS")
	BMC_ADDRESS = os.Getenv("BERLIN_BMC_ADDRESS")

	// check if any of the env vars are empty, panic if so
	if PASSWORD == "" || BTP_ADDRESS_LINE == "" || XCALL_ADDRESS == "" || BMC_ADDRESS == "" {
		panic("One or more env vars are empty. Please check your .env file.")
	}

	// split the BTP_ADDRESS_LINE string at commas, MOVE TO CONFIG
	BTP_ADDRESSES_TO_TRACK = strings.Split(BTP_ADDRESS_LINE, ",")

	if config.TestMode {
		fmt.Println("")
		for adr := range BTP_ADDRESSES_TO_TRACK {
			fmt.Println("BTP_ADDRESS_TO_TRACK:", BTP_ADDRESSES_TO_TRACK[adr])
		}
		fmt.Println("")
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
		fmt.Println("Connecting to Berlin...")

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
