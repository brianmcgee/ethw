package wallet

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/sha3"
)

type Wallet struct {
	Alias      string `json:"alias"`
	Address    string `json:"address"`
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
}

// NewWallet creates a new WalletInfo from the given seed and alias.
func NewWallet(seed []byte, alias string) (*Wallet, error) {
	buf, err := generateHash(seed)
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.ToECDSA(buf)
	if err != nil {
		return nil, fmt.Errorf("failed to create private key: %w", err)
	}

	publicKeyECDSA, err := extractPublicKey(privateKey)
	if err != nil {
		return nil, err
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return &Wallet{
		Alias:      alias,
		Address:    address,
		PrivateKey: hex.EncodeToString(crypto.FromECDSA(privateKey)),
		PublicKey:  hex.EncodeToString(crypto.FromECDSAPub(publicKeyECDSA)),
	}, nil
}

func generateHash(seed []byte) ([]byte, error) {
	hash := sha3.NewLegacyKeccak256()
	_, err := hash.Write(seed)
	if err != nil {
		return nil, err
	}
	return hash.Sum(nil), nil
}

func extractPublicKey(privateKey *ecdsa.PrivateKey) (*ecdsa.PublicKey, error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key to ECDSA")
	}
	return publicKeyECDSA, nil
}

type Seed struct {
	Alias string
	Data  []byte
}

func (s *Seed) UnmarshalText(text []byte) error {
	return s.processSeedString(strings.TrimSpace(string(text)))
}

func (s *Seed) processSeedString(seedStr string) error {
	splitData := strings.SplitN(seedStr, ":", 2)
	if len(splitData) == 2 {
		s.Alias = strings.TrimSpace(splitData[0])
		seedStr = strings.TrimSpace(splitData[1])
	}

	log.Debug("Current word list", bip39.GetWordList())

	// TODO: Check why IsMnemonicValid doesn't pass correctly
	// if !bip39.IsMnemonicValid(seedStr) {
	// 	return fmt.Errorf("Invalid mnemonic string: %s", seedStr)
	// }

	s.Data = bip39.NewSeed(seedStr, "")

	return nil
}
