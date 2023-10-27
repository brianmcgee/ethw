package keystore

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	k "github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// keystore encapsulates a keystore directory and the underlying Ethereum keystore.
type KeystoreWrapper struct {
	ks  *k.KeyStore
	dir string
}

// NewKeyStore initializes a new Ethereum keystore and the directory where it's stored.
func NewKeyStore(dir string) *KeystoreWrapper {
	ks := k.NewKeyStore(dir, k.StandardScryptN, k.StandardScryptP)
	return &KeystoreWrapper{
		ks:  ks,
		dir: dir,
	}
}

// ImportPrivateKey imports a private key into the keystore, optionally overwriting an existing account.
func (kst *KeystoreWrapper) ImportPrivateKey(privateKeyHex string, password string, overwrite bool) error {
	// Decode the private key.
	key, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return fmt.Errorf("failed to decode private key: %w", err)
	}

	address := crypto.PubkeyToAddress(key.PublicKey)
	if kst.ks.HasAddress(address) && !overwrite {
		return fmt.Errorf("address already exists")
	}

	if kst.ks.HasAddress(address) && overwrite {
		if err := kst.UnsafeDeleteAccount(address); err != nil {
			return fmt.Errorf("failed to delete existing account: %w", err)
		}
	}

	// Import the new account into the keystore.
	if _, err := kst.ks.ImportECDSA(key, password); err != nil {
		return fmt.Errorf("failed to import private key: %w", err)
	}

	return nil
}

// UnsafeDeleteAccount deletes an Ethereum account without requiring its password.
// TODO: Probably will be removed but leaving it for now
func (kst *KeystoreWrapper) UnsafeDeleteAccount(address common.Address) error {
	// Check if the address exists in the keystore.
	if !kst.ks.HasAddress(address) {
		return fmt.Errorf("address does not exist")
	}

	// Read filenames in the keystore directory.
	files, err := os.ReadDir(kst.dir)
	if err != nil {
		return fmt.Errorf("failed to read keystore directory: %w", err)
	}

	// Search for the file corresponding to the address and delete it.
	for _, file := range files {
		if strings.Contains(file.Name(), strings.ToLower(address.Hex()[2:])) {
			err := os.Remove(filepath.Join(kst.dir, file.Name()))
			if err != nil {
				return fmt.Errorf("failed to delete keystore file: %w", err)
			}
			kst.ks.Wallets() // Calling this, will refresh internally the list of wallets (inneficient but works)
			return nil
		}
	}

	return fmt.Errorf("keystore file for the address %s not found", strings.ToLower(address.Hex()))
}

// Accounts returns all key files present in the directory.
func (kst *KeystoreWrapper) Accounts() []accounts.Account {
	return kst.ks.Accounts()
}
