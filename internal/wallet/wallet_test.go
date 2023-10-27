package wallet

import (
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

// TestNewWallet checks if NewWallet correctly creates a wallet from a seed and alias.
func TestNewWallet(t *testing.T) {
	seed := []byte("testseed")
	alias := "testalias"

	wallet, err := NewWallet(seed, alias)

	assert.Nil(t, err)
	assert.Equal(t, alias, wallet.Alias)
	assert.Equal(t, 42, len(wallet.Address))
	assert.Equal(t, 64, len(wallet.PrivateKey))
	assert.Equal(t, 130, len(wallet.PublicKey))
}

// TestGenerateHash checks the generateHash function.
func TestGenerateHash(t *testing.T) {
	seed := []byte("test")
	hash, err := generateHash(seed)

	assert.Nil(t, err)
	assert.Equal(t, 32, len(hash))
}

// TestExtractPublicKey checks if it correctly extracts a public key from a private key.
func TestExtractPublicKey(t *testing.T) {
	seed := []byte("test")
	hash, _ := generateHash(seed)
	privateKey, _ := crypto.ToECDSA(hash)

	publicKey, err := extractPublicKey(privateKey)

	assert.Nil(t, err)
	assert.NotNil(t, publicKey)
}

// TestWalletData_UnmarshalText_ValidData tests the successful unmarshalling of WalletData.
func TestWalletData_UnmarshalText_ValidData(t *testing.T) {
	text := []byte("seed=myseed;password=mypassword")
	wd := WalletData{}

	err := wd.UnmarshalText(text)

	assert.Nil(t, err)
	assert.Equal(t, "myseed", wd.Seed)
	assert.Equal(t, "mypassword", wd.Password)
}

// TestWalletData_UnmarshalText_MissingSeed tests unmarshalling with a missing seed.
func TestWalletData_UnmarshalText_MissingSeed(t *testing.T) {
	text := []byte("password=mypassword")
	wd := WalletData{}

	err := wd.UnmarshalText(text)

	assert.Equal(t, ErrInvalidWalletDataFormat, err)
}

// TestWalletData_UnmarshalText_MissingPassword tests unmarshalling with a missing password.
func TestWalletData_UnmarshalText_MissingPassword(t *testing.T) {
	text := []byte("seed=myseed")
	wd := WalletData{}

	err := wd.UnmarshalText(text)

	assert.Nil(t, err)
	assert.Equal(t, "myseed", wd.Seed)
	assert.Empty(t, wd.Password)
}
