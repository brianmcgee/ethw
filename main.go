package main

import (
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/alecthomas/kong"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/pelletier/go-toml"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/sha3"
)

var cli struct {
	Seeds      []string `arg:"" optional:"" type:"string" help:"Deterministic seeds or BIP-39 mnemonics to generate keys."`
	Json       bool     `optional:"" short:"j" help:"Output results in JSON format"`
	ConfigFile string   `optional:"" short:"c" type:"path" help:"Path to TOML config file containing seeds"`
}

type Config struct {
	Seeds []string `toml:"seeds"`
}

type WalletInfo struct {
	Alias           string `json:"alias"`
	PrivateKey      string `json:"private_key"`
	PublicKey       string `json:"public_key"`
	EthereumAddress string `json:"ethereum_address"`
}

func loadSeeds() ([]string, error) {
	if cli.ConfigFile != "" {
		return readTomlConfig(cli.ConfigFile)
	}
	return cli.Seeds, nil
}

func readTomlConfig(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening config file: %w", err)
	}
	defer file.Close()

	config := &Config{}
	if err := toml.NewDecoder(file).Decode(config); err != nil {
		return nil, fmt.Errorf("error decoding TOML: %w", err)
	}

	return config.Seeds, nil
}

func processSeeds(seeds []string) ([]WalletInfo, error) {
	var walletInfos []WalletInfo

	for i, seedStr := range seeds {
		alias, seed, err := processSingleSeed(seedStr)
		if err != nil {
			log.Printf("Error processing seed: %v", err)
			continue
		}

		walletInfo, err := generateWallet(seed)
		if err != nil {
			log.Printf("Error generating wallet: %v", err)
			continue
		}

		if alias != "" {
			walletInfo.Alias = alias
		} else {
			walletInfo.Alias = fmt.Sprintf("Wallet %d", i+1)
		}

		walletInfos = append(walletInfos, walletInfo)
	}
	return walletInfos, nil
}

func processSingleSeed(seedStr string) (string, []byte, error) {
	var alias string

	splitData := strings.SplitN(seedStr, "|", 2)
	if len(splitData) == 2 {
		alias = strings.TrimSpace(splitData[0])
		seedStr = strings.TrimSpace(splitData[1])
	}

	if bip39.IsMnemonicValid(seedStr) {
		seed, err := bip39.NewSeedWithErrorChecking(seedStr, "")
		if err != nil {
			return "", nil, err
		}
		return alias, seed, nil
	}

	return alias, []byte(seedStr), nil
}

func generateWallet(seed []byte) (WalletInfo, error) {
	hash := sha3.NewLegacyKeccak256()
	hash.Write(seed)
	buf := hash.Sum(nil)

	privateKey, err := crypto.ToECDSA(buf)
	if err != nil {
		return WalletInfo{}, fmt.Errorf("failed to create private key: %w", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return WalletInfo{}, fmt.Errorf("error casting public key to ECDSA")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return WalletInfo{
		PrivateKey:      hex.EncodeToString(crypto.FromECDSA(privateKey)),
		PublicKey:       hex.EncodeToString(crypto.FromECDSAPub(publicKeyECDSA)),
		EthereumAddress: address,
	}, nil
}

func renderJsonOutput(walletInfos []WalletInfo) {
	jsonOutput, err := json.Marshal(walletInfos)
	if err != nil {
		log.Fatalf("JSON Marshalling failed: %v", err)
	}
	fmt.Println(string(jsonOutput))
}

func renderTableOutput(walletInfos []WalletInfo) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Wallet", "Ethereum Address", "Private Key", "Public Key"})
	for i, walletInfo := range walletInfos {
		alias := walletInfo.Alias
		if alias == "" {
			alias = fmt.Sprintf("Wallet %d", i+1)
		}
		t.AppendRow([]interface{}{alias, walletInfo.EthereumAddress, walletInfo.PrivateKey, walletInfo.PublicKey})
	}
	t.Render()
}

func main() {
	kong.Parse(&cli)

	seeds, err := loadSeeds()
	if err != nil {
		log.Fatalf("Failed to load seeds: %v", err)
	}

	walletInfos, err := processSeeds(seeds)
	if err != nil {
		log.Fatalf("Failed to process seeds: %v", err)
	}

	if cli.Json {
		renderJsonOutput(walletInfos)
	} else {
		renderTableOutput(walletInfos)
	}
}
