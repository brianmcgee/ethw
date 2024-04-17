package cmd

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/aldoborrero/ethw/internal/utils/output"
	"github.com/aldoborrero/ethw/internal/wallet"
	"github.com/alecthomas/kong"
	"github.com/charmbracelet/log"
	"github.com/tyler-smith/go-bip39"
)

type walletCreateCmd struct {
	Mnemonic []MnemonicData `arg:"" type:"custom" help:"Deterministic BIP-39 mnemonics to generate wallets"`
}

func (cmd *walletCreateCmd) Run(ctx *kong.Context) error {
	walletInfos, errs := processMnemonics(cmd.Mnemonic)
	if len(errs) > 0 {
		for _, err := range errs {
			log.Printf("%v", err)
		}
		return fmt.Errorf("there were errors processing seeds")
	}

	var writer output.WalletOutputWriter
	switch Cli.OutputFormat {
	case "json":
		writer = output.WalletJSONOutputWriter{}
	case "csv":
		writer = output.WalletCSVOutputWriter{}
	case "table":
		writer = output.WalletTableOutputWriter{}
	default:
		writer = output.WalletTextOuputWriter{}
	}

	if err := writer.WriteCreateOutput(walletInfos); err != nil {
		return fmt.Errorf("failed to generate output: %w", err)
	}

	return nil
}

func processMnemonics(mnemonics []MnemonicData) ([]*wallet.Wallet, []error) {
	var wallets []*wallet.Wallet
	var errors []error

	for i, mnemonic := range mnemonics {
		walletInfo, err := wallet.NewWallet(mnemonic.Mnemonic, mnemonic.Alias)
		if err != nil {
			errors = append(errors, fmt.Errorf("error generating wallet for seed %d: %w", i, err))
			continue
		}

		wallets = append(wallets, walletInfo)
	}

	return wallets, errors
}

var (
	// errInvalidSeedFormat represents an error when the seed string is invalid.
	errInvalidSeedFormat = errors.New("invalid seed format")

	// errInvalidSeedMnemonic represents an error when the mnemonic list is invalid.
	errInvalidSeedMnemonic = errors.New("invalid mnemonic format")

	// seedDataRegex is a compiled regular expression for parsing seed strings.
	seedDataRegex = regexp.MustCompile(`(?:seed=([^;]*);?)(?:alias=([^;]*);?)?`)
)

// MnemonicData represents the seed information for a wallet.
type MnemonicData struct {
	Alias    string
	Mnemonic string
}

// UnmarshalText unmarshals the SeedData from text
func (sd *MnemonicData) UnmarshalText(raw []byte) error {
	matches := seedDataRegex.FindStringSubmatch(string(raw))

	// Check for invalid format or missing seed
	if matches == nil || matches[1] == "" {
		log.Debugf("No match found for the current seed")
		return errInvalidSeedFormat
	}

	rawMnemonic := strings.TrimSpace(matches[1])
	log.Debugf("Raw mnemonic: %s", rawMnemonic)

	if !bip39.IsMnemonicValid(rawMnemonic) {
		log.Debugf("Invalid nmnemonic passed")
		return errInvalidSeedMnemonic
	}

	sd.Mnemonic = rawMnemonic
	sd.Alias = strings.TrimSpace(matches[2])

	log.Debugf("Decoded SeedData: %+v", sd)

	return nil
}
