package keystore

import (
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

// CreateKeyStore generates a new keystore at the specified directory
func CreateKeyStore(outputDir string) (*keystore.KeyStore, error) {
	ks := keystore.NewKeyStore(outputDir, keystore.StandardScryptN, keystore.StandardScryptP)
	return ks, nil
}

// ImportPrivateKey imports a private key into the given keystore
func ImportPrivateKey(ks *keystore.KeyStore, privateKeyHex string, password string) error {
	key, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return err
	}

	_, err = ks.ImportECDSA(key, password)
	return err
}
