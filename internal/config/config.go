// config/config.go
package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

var (
	TestMode               bool
	BerlinTestBlockHeight  int64
	SepoliaTestBlockHeight int64
	BTP_ADDRESSES_TO_TRACK []string
	XCALL_ADDRESS          string
	BERLIN_BMC_ADDRESS     string
	RELAY_ADDRESS          string
	SEPOLIA_BMC_ADDRESS    string
	SEPOLIA_XCALL_ADDRESS  string
)

// Should also include a config option to allow for signing txs with a keystore file.
// Atm testmode only prints to console instead of signing a execute call tx

func init() {
	// test for test flag
	args := os.Args
	if len(args) == 2 {
		if args[1] == "test" {
			TestMode = true
			fmt.Printf("\nRunning in test mode. Be sure that the blocks to check are set correctly in .env\n\n")
		} else {
			TestMode = false
		}
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	SEPOLIA_BMC_ADDRESS = os.Getenv("SEPOLIA_BMC_ADDRESS")
	if SEPOLIA_BMC_ADDRESS == "" {
		panic("SEPOLIA_BMC_ADDRESS is empty. Please check your .env file.")
	}

	SEPOLIA_XCALL_ADDRESS = os.Getenv("SEPOLIA_XCALL_ADDRESS")
	if SEPOLIA_XCALL_ADDRESS == "" {
		panic("SEPOLIA_XCALL_ADDRESS is empty. Please check your .env file.")
	}

	BTP_ADDRESS_LINE := os.Getenv("BTP_ADDRESS_TO_TRACK")
	XCALL_ADDRESS = os.Getenv("BERLIN_XCALL_ADDRESS")
	BERLIN_BMC_ADDRESS = os.Getenv("BERLIN_BMC_ADDRESS")

	// check if any of the env vars are empty, panic if so
	if BTP_ADDRESS_LINE == "" || XCALL_ADDRESS == "" || BERLIN_BMC_ADDRESS == "" {
		panic("One or more env vars are empty. Please check your .env file.")
	}

	// split the BTP_ADDRESS_LINE string at commas, MOVE TO CONFIG
	BTP_ADDRESSES_TO_TRACK = strings.Split(BTP_ADDRESS_LINE, ",")

	if TestMode {
		fmt.Println("")
		for adr := range BTP_ADDRESSES_TO_TRACK {
			fmt.Println("BTP_ADDRESS_TO_TRACK:", BTP_ADDRESSES_TO_TRACK[adr])
		}
		fmt.Println("")
	}

	if !TestMode {
		return
	}

	// if testmode is true get the block to test from the env file
	BerlinTestBlock := os.Getenv("BERLIN_TEST_BLOCKHEIGHT")
	if BerlinTestBlock != "" {
		BerlinTestBlockHeight, _ = strconv.ParseInt(BerlinTestBlock, 10, 64)
		fmt.Println("Berlin Test BlockHeight: ", BerlinTestBlockHeight)
	} else {
		panic("BerlinTestBlockHeight is empty. Please check your .env file.")
	}

	// and for sepolia
	SepoliaTestBlock := os.Getenv("SEPOLIA_TEST_BLOCKHEIGHT")
	if SepoliaTestBlock != "" {
		SepoliaTestBlockHeight, _ = strconv.ParseInt(SepoliaTestBlock, 10, 64)
		fmt.Println("Sepolia Test BlockHeight: ", SepoliaTestBlockHeight)
	} else {
		panic("SepoliaTestBlockHeight is empty. Please check your .env file.")
	}

}
