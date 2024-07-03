package wallet

import (
	"fmt"

	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

type Wallet struct {
	Alias      string `json:"alias"`
	Address    string `json:"address"`
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
}

// NewWallet creates a new Wallet from the given mnemonic, alias, and an optional custom derivation path.
// It uses the provided account derived from the standard Ethereum derivation path if no custom path is provided.
func NewWallet(mnemonic, alias string, customDerivationPath ...string) (*Wallet, error) {
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return nil, fmt.Errorf("failed to create wallet from mnemonic: %w", err)
	}

	// Use the default base derivation path unless a custom path is provided
	derivationPath := hdwallet.DefaultBaseDerivationPath.String()
	if len(customDerivationPath) > 0 && customDerivationPath[0] != "" {
		derivationPath = customDerivationPath[0]
	}

	path, err := hdwallet.ParseDerivationPath(derivationPath)
	if err != nil {
		return nil, fmt.Errorf("failed to parse derivation path: %w", err)
	}

	account, err := wallet.Derive(path, true)
	if err != nil {
		return nil, fmt.Errorf("failed to derive account: %w", err)
	}

	privateKeyHex, err := wallet.PrivateKeyHex(account)
	if err != nil {
		return nil, fmt.Errorf("failed to get private key hex: %w", err)
	}

	publicKeyHex, err := wallet.PublicKeyHex(account)
	if err != nil {
		return nil, fmt.Errorf("failed to get public key hex: %w", err)
	}

	return &Wallet{
		Alias:      alias,
		Address:    account.Address.Hex(),
		PrivateKey: privateKeyHex,
		PublicKey:  publicKeyHex,
	}, nil
}
