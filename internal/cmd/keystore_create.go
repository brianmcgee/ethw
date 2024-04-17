package cmd

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/aldoborrero/ethw/internal/keystore"
	"github.com/aldoborrero/ethw/internal/utils/output"
	"github.com/aldoborrero/ethw/internal/wallet"
	"github.com/alecthomas/kong"
	"github.com/charmbracelet/log"
	"github.com/tyler-smith/go-bip39"
)

type keystoreCreateCmd struct {
	Wallets     []WalletData `arg:"" type:"custom" help:"List of 'seed' and 'password' to generate wallets"`
	Overwrite   bool         `flag:"" optional:"" help:"Overwrite wallet creation if exists one in the keystore"`
	KeystoreDir string       `flag:"" optional:"" type:"path" default:"./keystore" help:"Directory to save the keystore file"`
}

func (cmd *keystoreCreateCmd) Run() error {
	absKeystoreDir := kong.ExpandPath(cmd.KeystoreDir)

	if cmd.Overwrite {
		if err := os.RemoveAll(absKeystoreDir); err != nil {
			log.Error("Failed to remove keystore directory: ", err)
			return err
		}
	}

	ks := keystore.NewKeyStore(absKeystoreDir)

	for i, walletData := range cmd.Wallets {
		if err := cmd.createWallet(walletData, i, ks); err != nil {
			log.Error(err.Error())
			return err
		}
	}

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

	if err := writer.WriteCreateOutput(*ks); err != nil {
		return fmt.Errorf("failed to generate output: %w", err)
	}

	return nil
}

func (cmd *keystoreCreateCmd) createWallet(walletData WalletData, index int, ks *keystore.KeystoreWrapper) error {
	mnemonic := walletData.Mnemonic
	password := walletData.Password

	walletInstance, err := wallet.NewWallet(mnemonic, "")
	if err != nil {
		log.Errorf("Failed to generate wallet %d from seed: %v", index+1, err)
		return err
	}

	log.Infof("Creating wallet %d with address %s", index+1, walletInstance.Address)
	if err := ks.ImportPrivateKey(walletInstance.PrivateKey, password, false); err != nil {
		log.Errorf("Failed to import private key into keystore for wallet %d: %v", index+1, err)
		return err
	}
	return nil
}

var (
	reWalletData                     = regexp.MustCompile(`(?:seed=([^;]*))(?:;password=([^;]*))?(?:;path=([^;]*))?`)
	errInvalidWalletDataFormat       = errors.New("invalid wallet format")
	errInvalidWalletDataMnemonicSeed = errors.New("invalid mnemonic format")
)

type WalletData struct {
	Mnemonic       string
	Password       string
	DerivationPath string
}

func (wd *WalletData) UnmarshalText(raw []byte) error {
	matches := reWalletData.FindStringSubmatch(string(raw))
	if matches == nil || len(matches[1]) == 0 {
		return errInvalidWalletDataFormat
	}

	mnemonic, password, path := strings.TrimSpace(matches[1]), strings.TrimSpace(matches[2]), strings.TrimSpace(matches[3])

	if !bip39.IsMnemonicValid(mnemonic) {
		return errInvalidWalletDataMnemonicSeed
	}

	*wd = WalletData{
		Mnemonic:       mnemonic,
		Password:       password,
		DerivationPath: path,
	}

	return nil
}
