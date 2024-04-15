package output

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aldoborrero/ethw/internal/wallet"
	"github.com/jedib0t/go-pretty/v6/table"
)

// WalletOutputWriter is an interface for writing wallet information to different output formats.
type WalletOutputWriter interface {
	WriteCreateOutput([]*wallet.Wallet) error
}

type WalletTextOuputWriter struct{}

// WriteCreateOutput writes the details of the wallets in a clear, readable, text format.
func (w WalletTextOuputWriter) WriteCreateOutput(wallets []*wallet.Wallet) error {
	if len(wallets) == 0 {
		fmt.Println("No wallets created")
		return nil
	}

	fmt.Println("Wallets Information:")
	for i, walletInfo := range wallets {
		fmt.Printf("  Wallet #%d:\n", i+1)
		fmt.Printf("    Alias: %s\n", walletInfo.Alias)
		fmt.Printf("    Address: %s\n", walletInfo.Address)
		fmt.Printf("    Private Key: %s\n", walletInfo.PrivateKey)
		fmt.Printf("    Public Key: %s\n\n", walletInfo.PublicKey)
	}

	return nil
}

// WalletTableOutputWriter is a type that implements the OutputWriter interface for table-formatted data.
type WalletTableOutputWriter struct{}

// WriteCreateOutput writes the details of the wallets to standard output in table format.
// Returns nil as it does not encounter any errors in the process.
func (t WalletTableOutputWriter) WriteCreateOutput(walletInfos []*wallet.Wallet) error {
	tw := table.NewWriter()
	tw.SetOutputMirror(os.Stdout)
	tw.AppendHeader(table.Row{"#", "Alias", "Address", "Private Key", "Public Key"})
	for i, walletInfo := range walletInfos {
		tw.AppendRow([]interface{}{i + 1, walletInfo.Alias, walletInfo.Address, walletInfo.PrivateKey, walletInfo.PublicKey})
	}
	tw.Render()
	return nil
}

// WalletJSONOutputWriter is a type that implements the OutputWriter interface for JSON-formatted data.
type WalletJSONOutputWriter struct{}

// WriteWalletOutput writes the details of the wallets to standard output in JSON format.
// Returns an error if JSON marshaling fails.
func (j WalletJSONOutputWriter) WriteCreateOutput(walletInfos []*wallet.Wallet) error {
	jsonOutput, err := json.Marshal(walletInfos)
	if err != nil {
		return err
	}
	fmt.Println(string(jsonOutput))
	return nil
}

// WalletCSVOutputWriter writes wallet information in CSV format.
type WalletCSVOutputWriter struct{}

// WriteCreateOutput writes the details of the wallets in CSV format to standard output.
func (w WalletCSVOutputWriter) WriteCreateOutput(wallets []*wallet.Wallet) error {
	csvWriter := csv.NewWriter(os.Stdout)
	defer csvWriter.Flush()

	header := []string{"#", "Alias", "Address", "Private Key", "Public Key"}
	if err := csvWriter.Write(header); err != nil {
		return fmt.Errorf("writing CSV header: %w", err)
	}

	for i, walletInfo := range wallets {
		record := []string{
			fmt.Sprintf("%d", i+1),
			walletInfo.Alias,
			walletInfo.Address,
			walletInfo.PrivateKey,
			walletInfo.PublicKey,
		}
		if err := csvWriter.Write(record); err != nil {
			return fmt.Errorf("writing CSV record for wallet #%d: %w", i+1, err)
		}
	}

	return nil
}
