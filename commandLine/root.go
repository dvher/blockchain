package commandLine

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "blockchain.exe",
	Short: "blockchain is a CLI for blockchain",
	Long:  `blockchain is a CLI for blockchain`,
	Run: func(cmd *cobra.Command, args []string) {
	},
	ValidArgs: []string{"getbalance", "createblockchain", "printchain", "searchblock", "send", "createwallet", "listaddresses"},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Panic(err)
	}
}

func init() {

	if len(os.Args) == 1 {
		rootCmd.SetArgs([]string{"help"})
	}

	var cmdGetBalance = &cobra.Command{
		Use:   "getbalance [address]",
		Short: "Gets the balance linked to the address provided",
		Args:  cobra.ExactArgs(1),
		Run:   getBalance,
	}

	var cmdCreateBlockchain = &cobra.Command{
		Use:   "createblockchain [address]",
		Short: "Creates a blockchain with the given address as genesis",
		Args:  cobra.ExactArgs(1),
		Run:   createBlockChain,
	}

	var cmdPrintChain = &cobra.Command{
		Use:   "printchain",
		Short: "Print the current chain",
		Args:  cobra.ExactArgs(0),
		Run:   printChain,
	}

	var cmdSearchBlock = &cobra.Command{
		Use:   "searchblock [BlockHash]",
		Short: "Search the given block by hash",
		Args:  cobra.ExactArgs(1),
		Run:   searchBlockByHash,
	}

	var cmdSend = &cobra.Command{
		Use:   "send",
		Short: "Send money from an account to another given both addresses and amount",
		Args:  cobra.ExactArgs(0),
		Run:   send,
	}

	var cmdCreateWallet = &cobra.Command{
		Use:   "createwallet",
		Short: "Creates a new wallet and print the info",
		Args:  cobra.ExactArgs(0),
		Run:   createWallet,
	}

	var cmdListAddresses = &cobra.Command{
		Use:   "listaddresses",
		Short: "Lists all wallet addresses in the chain",
		Args:  cobra.ExactArgs(0),
		Run:   listAddresses,
	}

	cmdSend.Flags().StringVarP(&sendFrom, "from", "f", "", "Address from which to send coins")
	cmdSend.Flags().StringVarP(&sendTo, "to", "t", "", "Address to send coins to")
	cmdSend.Flags().IntVarP(&amount, "amount", "a", 0, "Amount to be sent")
	cmdSend.MarkFlagRequired("from")
	cmdSend.MarkFlagRequired("to")
	cmdSend.MarkFlagRequired("amount")

	rootCmd.AddCommand(cmdGetBalance)
	rootCmd.AddCommand(cmdCreateBlockchain)
	rootCmd.AddCommand(cmdPrintChain)
	rootCmd.AddCommand(cmdSearchBlock)
	rootCmd.AddCommand(cmdSend)
	rootCmd.AddCommand(cmdCreateWallet)
	rootCmd.AddCommand(cmdListAddresses)

}
