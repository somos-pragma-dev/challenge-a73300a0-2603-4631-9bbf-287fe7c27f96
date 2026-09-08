package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"internal/services"
)

func BalanceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "balance",
		Short: "Check account balance",
		RunE: func(cmd *cobra.Command, args []string) error {
			balance, err := services.GetBalance()
			if err!= nil {
				return err
			}
			fmt.Printf("Current balance: %f\n", balance)
			return nil
		},
	}
	return cmd
}