package cmd

import (
	"fmt"

	"github.com/aldoborrero/ethw/internal/keystore"
	"github.com/aldoborrero/ethw/internal/wallet"
	"github.com/alecthomas/kong"
	"github.com/charmbracelet/log"
)

type keystoreCreateCmd struct {
	PrivateKey *string      `flag:"" optional:"" type:"string" help:"Private key to store in keystore" xor:"Seed"`
	Seed       *wallet.Seed `flag:"" optional:"" type:"custom" help:"Seed to generate a private key" xor:"PrivateKey"`
	Password   string       `flag:"" optional:"" type:"string" default:"" help:"Passphrase to encrypt the keystore"`
	OutputDir  string       `arg:"" optional:"" type:"path" default:"./keystore" help:"Directory to save the keystore file"`
}

func (cmd *keystoreCreateCmd) Validate() error {
	if cmd.PrivateKey != nil && cmd.Seed != nil && string(cmd.Seed.Data) != "" {
		return fmt.Errorf("Either private-key or seed must be provided, but not both.")
	}
	if cmd.PrivateKey == nil && (cmd.Seed == nil || string(cmd.Seed.Data) == "") {
		return fmt.Errorf("Either private-key or seed must be provided.")
	}
	return nil
}

func (cmd *keystoreCreateCmd) Run() error {
	absOutputDir := kong.ExpandPath(cmd.OutputDir)

	log.Infof("Creating keystore at: %s", absOutputDir)
	ks, err := keystore.CreateKeyStore(absOutputDir)
	if err != nil {
		log.Errorf("Failed to create keystore: %v", err)
		return err
	}

	if cmd.Seed != nil && string(cmd.Seed.Data) != "" {
		wallet, err := wallet.NewWallet(cmd.Seed.Data, "")
		if err != nil {
			log.Errorf("Failed to generate wallet from seed: %v", err)
			return err
		}
		cmd.PrivateKey = &wallet.PrivateKey
	}

	log.Info("Keystore created. Importing private key...")
	if err := keystore.ImportPrivateKey(ks, *cmd.PrivateKey, cmd.Password); err != nil {
		log.Errorf("Failed to import private key into keystore: %v", err)
		return err
	}

	log.Info("Private key successfully imported into keystore.")

	return nil
}
