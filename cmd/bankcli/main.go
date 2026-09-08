package main

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"internal/commands"
)

func main() {
	rootCmd := cobra.Command{
		Use: "bankcli",
		Short: "Bank CLI for managing digital banking operations",
	}
	rootCmd.AddCommand(commands.BalanceCmd())
	rootCmd.AddCommand(commands.TransferCmd())
	if err := rootCmd.Execute(); err!= nil {
		fmt.Println(err)
		os.Exit(1)
	}
}