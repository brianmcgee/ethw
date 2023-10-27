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
	WriteCreateOutput([]*wallet.Wallet) error
}

type WalletTextOuputWriter struct{}

func (w WalletTextOuputWriter) WriteCreateOutput([]*wallet.Wallet) error {
	return nil
}

// TableOutputWriter is a type that implements the OutputWriter interface for table-formatted data.
type TableOutputWriter struct{}

// WriteCreateOutput writes the details of the wallets to standard output in table format.
// Returns nil as it does not encounter any errors in the process.
func (t TableOutputWriter) WriteCreateOutput(walletInfos []*wallet.Wallet) error {
	tw := table.NewWriter()
	tw.SetOutputMirror(os.Stdout)
	tw.AppendHeader(table.Row{"#", "Alias", "Address", "Private Key", "Public Key"})
	for i, walletInfo := range walletInfos {
		alias := walletInfo.Alias
		tw.AppendRow([]interface{}{i + 1, alias, walletInfo.Address, walletInfo.PrivateKey, walletInfo.PublicKey})
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
