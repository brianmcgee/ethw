package output

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aldoborrero/ethw/internal/wallet"
	"github.com/jedib0t/go-pretty/v6/table"
)

// WalletOutputWriter is an interface for writing wallet information to different output formats.
type WalletOutputWriter interface {
	WriteOutput([]*wallet.Wallet) error
}

// TableDataWriter is a type that implements the OutputWriter interface for table-formatted data.
type TableDataWriter struct{}

// WriteWalletOutput writes the details of the wallets to standard output in table format.
// Returns nil as it does not encounter any errors in the process.
func (t TableDataWriter) WriteOutput(walletInfos []*wallet.Wallet) error {
	tw := table.NewWriter()
	tw.SetOutputMirror(os.Stdout)
	tw.AppendHeader(table.Row{"Wallet", "Address", "Private Key", "Public Key"})
	for i, walletInfo := range walletInfos {
		alias := walletInfo.Alias
		if alias == "" {
			alias = fmt.Sprintf("Wallet %d", i+1)
		}
		tw.AppendRow([]interface{}{alias, walletInfo.Address, walletInfo.PrivateKey, walletInfo.PublicKey})
	}
	tw.Render()
	return nil
}

// JSONDataWriter is a type that implements the OutputWriter interface for JSON-formatted data.
type JSONDataWriter struct{}

// WriteWalletOutput writes the details of the wallets to standard output in JSON format.
// Returns an error if JSON marshaling fails.
func (j JSONDataWriter) WriteOutput(walletInfos []*wallet.Wallet) error {
	jsonOutput, err := json.Marshal(walletInfos)
	if err != nil {
		return err
	}
	fmt.Println(string(jsonOutput))
	return nil
}
