package cmd

import (
	"fmt"

	"github.com/aldoborrero/ethw/internal/keystore"
)

type keystoreListCmd struct {
	KeystorePath string `flag:"" optional:"" type:"path" default:"./keystore" help:"Directory where the keystore is located"`
}

func (cmd *keystoreListCmd) Run() error {
	// Initialize the keystore
	ks, err := keystore.Initialize(cmd.KeystorePath)
	if err != nil {
		return fmt.Errorf("failed to obtain the keystore")
	}

	// Fetch all accounts
	accounts := ks.Accounts()

	// Check if there are any accounts to display
	if len(accounts) == 0 {
		fmt.Println("No accounts found")
		return nil
	}

	// Display accounts
	fmt.Println("List of Accounts:")
	for i, account := range accounts {
		fmt.Printf("%d. Address: %s", i+1, account.Address.Hex())
	}

	return nil
}
