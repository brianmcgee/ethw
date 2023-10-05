package cmd

import (
	"fmt"
	"log"

	"github.com/aldoborrero/ethw/internal/utils"
	"github.com/aldoborrero/ethw/internal/wallet"
	"github.com/alecthomas/kong"
)

type walletCreateCmd struct {
	Seeds []wallet.Seed `flag:"" optional:"" type:"custom" help:"Deterministic seeds or BIP-39 mnemonics to generate keys"`
	Json  bool          `optional:"" short:"j" help:"Output results in JSON format"`
}

func processSeeds(seeds []wallet.Seed) ([]*wallet.Wallet, []error) {
	var walletInfos []*wallet.Wallet
	var errors []error

	for i, seed := range seeds {
		walletInfo, err := wallet.NewWallet(seed.Data, seed.Alias)
		if err != nil {
			errors = append(errors, fmt.Errorf("Error generating wallet for seed %d: %w", i, err))
			continue
		}

		if seed.Alias == "" {
			walletInfo.Alias = fmt.Sprintf("Wallet %d", i+1)
		}

		walletInfos = append(walletInfos, walletInfo)
	}

	return walletInfos, errors
}

func (cmd *walletCreateCmd) Run(ctx *kong.Context) error {
	walletInfos, errs := processSeeds(cmd.Seeds)
	if len(errs) > 0 {
		for _, err := range errs {
			log.Printf("%v", err)
		}
		return fmt.Errorf("There were errors processing seeds")
	}

	var writer utils.OutputWriter
	if cmd.Json {
		writer = utils.JSONDataWriter{}
	} else {
		writer = utils.TableDataWriter{}
	}

	if err := writer.WriteWalletOutput(walletInfos); err != nil {
		return fmt.Errorf("Failed to generate output: %w", err)
	}

	return nil
}
