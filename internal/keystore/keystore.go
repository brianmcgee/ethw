package keystore

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

// Initialize initializes a new keystore at the specified directory.
func Initialize(outputDir string) (*keystore.KeyStore, error) {
	return keystore.NewKeyStore(outputDir, keystore.StandardScryptN, keystore.StandardScryptP), nil
}

// ImportPrivateKey imports or overwrites a private key into the given keystore.
// If 'overwrite' is true, it will overwrite an existing account with the same address.
func ImportPrivateKey(ks *keystore.KeyStore, privateKeyHex string, password string, overwrite bool) error {
	key, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return fmt.Errorf("failed to decode private key: %v", err)
	}

	address := crypto.PubkeyToAddress(key.PublicKey)

	// Early exit if the address exists and the overwrite flag is not set.
	if ks.HasAddress(address) && !overwrite {
		return fmt.Errorf("address already exists")
	}

	// Delete the existing account if overwrite is true.
	if ks.HasAddress(address) && overwrite {
		account := accounts.Account{Address: address}
		if err := ks.Delete(account, password); err != nil {
			return fmt.Errorf("failed to delete existing account: %v", err)
		}
	}

	// Import into the ks the account
	if _, err := ks.ImportECDSA(key, password); err != nil {
		return fmt.Errorf("failed to import private key: %v", err)
	}

	return nil
}
