package cmd

import (
	"fmt"
	"os"

	"github.com/aldoborrero/ethw/internal/keystore"
	"github.com/aldoborrero/ethw/internal/utils/output"
	"github.com/aldoborrero/ethw/internal/wallet"
	"github.com/alecthomas/kong"
	"github.com/charmbracelet/log"
)

type keystoreCreateCmd struct {
	Wallets     []wallet.WalletData `arg:"" type:"custom" help:"List of 'seed' and 'password' to generate wallets"`
	Overwrite   bool                `flag:"" optional:"" help:"Overwrite wallet creation if exists one in the keystore"`
	KeystoreDir string              `flag:"" optional:"" type:"path" default:"./keystore" help:"Directory to save the keystore file"`
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

func (cmd *keystoreCreateCmd) createWallet(walletData wallet.WalletData, index int, ks *keystore.KeystoreWrapper) error {
	seed := walletData.Seed
	password := walletData.Password

	walletInstance, err := wallet.NewWallet([]byte(seed), "")
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
