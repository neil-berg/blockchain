package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/neil-berg/blockchain/blockchain"
)

// CLI is the command line interface shape
type CLI struct {
	Chain *blockchain.Blockchain
}

func (cli *CLI) printUsage() {
	fmt.Println("Error parsing CLI commands. \nCLI usage:")
	fmt.Println("\taddblock --data <some data>")
	fmt.Println("\tprintchain")
}

// Run starts the CLI
func (cli *CLI) Run() {
	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	addBlockData := addBlockCmd.String("data", "", "String of the block data")

	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	if len(os.Args) == 1 {
		cli.printUsage()
		return
	}

	switch os.Args[1] {
	case "addblock":
		addBlockCmd.Parse(os.Args[2:])
	case "printchain":
		printChainCmd.Parse(os.Args[2:])
	default:
		cli.printUsage()
		// os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}

func (cli *CLI) addBlock(data string) {
	err := cli.Chain.AddBlock(data)
	if err != nil {
		fmt.Println("Failed to add block")
	}
	fmt.Printf("Successfully added %s", data)
}

func (cli *CLI) printChain() {
	iterator := cli.Chain.GetNewIterator()

	for {
		block := iterator.Next()
		isGenesis := len(block.PrevHash) == 0
		if isGenesis {
			fmt.Println("======== GENESIS ========")
		} else {
			fmt.Println("======== BLOCK ========")
		}
		fmt.Printf("Timestamp:\t %v\n", block.Timestamp)
		fmt.Printf("Data:\t\t %s\n", block.Data)
		fmt.Printf("Hash:\t\t %x\n", block.Hash)
		fmt.Printf("Previous hash:\t %x\n", block.PrevHash)
		fmt.Printf("Nonce: \t\t %d\n", block.Nonce)

		if isGenesis {
			break
		}
	}
}
