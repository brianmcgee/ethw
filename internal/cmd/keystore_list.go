package cmd

import (
	"fmt"

	"github.com/aldoborrero/ethw/internal/keystore"
	"github.com/aldoborrero/ethw/internal/utils/output"
)

type keystoreListCmd struct {
	KeystoreDir string `flag:"" optional:"" type:"path" default:"./keystore" help:"Directory where the keystore is located"`
}

func (cmd *keystoreListCmd) Run() error {
	// Initialize the keystore
	ks := keystore.NewKeyStore(cmd.KeystoreDir)

	// Fetch all accounts (if any)
	accounts := ks.Accounts()

	// Prepare output writer
	var writer output.KeystoreOutputWriter
	switch Cli.OutputFormat {
	case "json":
		writer = output.KeystoreJSONOutputWriter{}
	case "csv":
		writer = output.KeystoreCSVOutputWriter{}
	case "table":
		writer = output.KeystoreTableOutputWriter{}
	default:
		writer = output.KeystoreTextOutputWriter{}
	}

	// Write result
	if err := writer.WriteListOutput(accounts); err != nil {
		return fmt.Errorf("failed to generate output: %w", err)
	}

	return nil
}
