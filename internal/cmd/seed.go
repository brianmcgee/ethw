package cmd

import (
	"fmt"

	"github.com/aldoborrero/ethw/internal/utils/output"
	"github.com/tyler-smith/go-bip39"
)

type seedCreateCmd struct {
	SeedPassword   string `flag:"" optional:"" default:"" short:"p" help:"Password for the seed"`
	MnemonicLength int    `flag:"" optional:"" default:"128" short:"m" enum:"128,160,192,224,256" help:"Entropy length for mnemonic. Can be 128, 160, 192, 224, or 256."`
	Json           bool   `optional:"" short:"j" help:"Output results in JSON format"`
}

func (cmd *seedCreateCmd) Run() error {
	// Generate a mnemonic for memorization or user-friendly seeds
	entropy, err := bip39.NewEntropy(cmd.MnemonicLength)
	if err != nil {
		return err
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return err
	}

	// Generate a Bip39 seed from the mnemonic and a password.
	seed := bip39.NewSeed(mnemonic, cmd.SeedPassword)

	var writer output.SeedOutputWriter
	if cmd.Json {
		writer = output.SeedJSONDataWriter{}
	} else {
		writer = output.SeedTableDataWriter{}
	}

	if err := writer.WriteOutput(mnemonic, seed); err != nil {
		return fmt.Errorf("Failed to generate output: %w", err)
	}

	return nil
}
