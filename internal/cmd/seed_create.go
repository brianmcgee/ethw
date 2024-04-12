package cmd

import (
	"fmt"

	"github.com/aldoborrero/ethw/internal/utils/output"
	"github.com/alecthomas/kong"
	"github.com/tyler-smith/go-bip39"
)

const (
	Entropy12Words = 128
	Entropy24Words = 256
)

type seedCreateCmd struct {
	SeedPassword string `flag:"" optional:"" default:"" short:"p" help:"Password for the seed"`
	Length       string `flag:"" optional:"" default:"12" short:"m" enum:"12,24" help:"Number of words in the mnemonic. Can be 12 or 24."`
	NumSeeds     int    `flag:"" optional:"" default:"1" short:"n" help:"Number of seeds to generate"`
}

func (c *seedCreateCmd) mapLengthToEntropy() (int, error) {
	switch c.Length {
	case "12":
		return Entropy12Words, nil
	case "24":
		return Entropy24Words, nil
	default:
		return 0, fmt.Errorf("Invalid mnemonic length: %s", c.Length)
	}
}

func (c *seedCreateCmd) Run(ctx *kong.Context) error {
	entropy, err := c.mapLengthToEntropy()
	if err != nil {
		return err
	}

	var mnemonics []string
	var seeds [][]byte

	for i := 0; i < c.NumSeeds; i++ {
		entropyBytes, err := bip39.NewEntropy(entropy)
		if err != nil {
			return err
		}

		mnemonic, err := bip39.NewMnemonic(entropyBytes)
		if err != nil {
			return err
		}

		seed := bip39.NewSeed(mnemonic, c.SeedPassword)

		mnemonics = append(mnemonics, mnemonic)
		seeds = append(seeds, seed)
	}

	var writer output.SeedOutputWriter
	switch Cli.OutputFormat {
	case "json":
		writer = output.SeedJSONOutputWriter{}
	case "csv":
		writer = output.SeedCSVOutputWriter{}
	case "table":
		writer = output.SeedTableOutputWriter{}
	default:
		writer = output.SeedTextOutputWriter{}
	}

	if err := writer.WriteOutput(mnemonics, seeds); err != nil {
		return fmt.Errorf("Failed to generate output: %w", err)
	}

	return nil
}
