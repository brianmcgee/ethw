package wallet

import (
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"regexp"

	"github.com/ethereum/go-ethereum/crypto"
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

// Define the regex pattern globally, so it is compiled only once.
var reWalletData = regexp.MustCompile(`(?:seed=([^;]+))?(?:;password=([^;]*))?`)

var ErrInvalidWalletDataFormat = errors.New("invalid wallet format")

// WalletData struct
type WalletData struct {
	Seed     string
	Password string
}

// Decode implements the MapperValue interface
func (wd *WalletData) UnmarshalText(text []byte) error {
	// Using pre-compiled regex for parsing
	matches := reWalletData.FindStringSubmatch(string(text))

	// Check for invalid format or missing seed
	if matches == nil || matches[1] == "" {
		return ErrInvalidWalletDataFormat
	}

	// Assign seed and password
	wd.Seed = matches[1]
	wd.Password = matches[2]

	return nil
}
