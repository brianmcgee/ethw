package cmd

import (
	"fmt"

	"github.com/aldoborrero/ethw/internal/keystore"
	"github.com/aldoborrero/ethw/internal/utils/output"
)

type keystoreListCmd struct {
	KeystoreDir string `flag:"" optional:"" type:"path" default:"./keystore" help:"Directory where the keystore is located"`
	Json        bool   `optional:"" short:"j" help:"Output results in JSON format"`
}

func (cmd *keystoreListCmd) Run() error {
	// Initialize the keystore
	ks := keystore.NewKeyStore(cmd.KeystoreDir)

	// Fetch all accounts (if any)
	accounts := ks.Accounts()

	// Prepare output writer
	var writer output.KeystoreOutputWriter
	if cmd.Json {
		writer = output.KeystoreJSONOutputWriter{}
	} else {
		writer = output.KeystoreTextOutputWriter{}
	}

	// Write result
	if err := writer.WriteListOutput(accounts); err != nil {
		return fmt.Errorf("failed to generate output: %w", err)
	}

	return nil
}
