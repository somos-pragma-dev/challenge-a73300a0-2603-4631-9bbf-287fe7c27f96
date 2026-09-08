package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"internal/services"
)

func TransferCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "transfer",
		Short: "Transfer funds between accounts",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Placeholder for transfer logic
			fmt.Println("Transfer command executed")
			return nil
		},
	}
	return cmd
}