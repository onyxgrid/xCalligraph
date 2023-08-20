// config/config.go
package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	TestMode              bool
	BerlinTestBlockHeight int64
    SepoliaTestBlockHeight int64
)

// Should also include a config option to allow for signing txs with a keystore file.
// Atm testmode only prints to console instead of signing a execute call tx

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}


	if os.Getenv("TEST_MODE") == "true" {
		TestMode = true
		fmt.Println("TestMode is true")
	} else {
		TestMode = false
	}

    if !TestMode {
        return
    }

    // if testmode is true get the block to test from the env file
	BerlinTestBlock := os.Getenv("BERLIN_TEST_BLOCKHEIGHT")
	if BerlinTestBlock != "" {
		BerlinTestBlockHeight, _ = strconv.ParseInt(BerlinTestBlock, 10, 64)
        fmt.Println("BerlinTestBlockHeight: ", BerlinTestBlockHeight)
	} else {
        panic("BerlinTestBlockHeight is empty. Please check your .env file.")
    }

    // and for sepolia
    SepoliaTestBlock := os.Getenv("SEPOLIA_TEST_BLOCKHEIGHT")
    if SepoliaTestBlock != "" {
        SepoliaTestBlockHeight, _ = strconv.ParseInt(SepoliaTestBlock, 10, 64)
        fmt.Println("SepoliaTestBlockHeight: ", SepoliaTestBlockHeight)
    } else {
        panic("SepoliaTestBlockHeight is empty. Please check your .env file.")
    }


}
