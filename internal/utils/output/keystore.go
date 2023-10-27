package output

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/jedib0t/go-pretty/v6/table"
)

// KeystoreOutputWriter is an interface for writing keystore information to different output formats.
type KeystoreOutputWriter interface {
	WriteCreateOutput(mnemonic string, seed []byte) error
	WriteListOutput(accounts []accounts.Account) error
}

// KeystoreTextOutputWriter writes keystore output in pure text format.
type KeystoreTextOutputWriter struct{}

func (w KeystoreTextOutputWriter) WriteCreateOutput(mnemonic string, seed []byte) error {
	return nil
}

func (w KeystoreTextOutputWriter) WriteListOutput(accounts []accounts.Account) error {
	if len(accounts) == 0 {
		fmt.Println("No accounts found")
		return nil
	}

	fmt.Println("List of Wallets:")
	for i, account := range accounts {
		fmt.Printf("%d.\t%s", i+1, account.Address.Hex())
	}

	return nil
}

// KeystoreTableOutputWriter writes keystore output in table format.
type KeystoreTableOutputWriter struct{}

func (w KeystoreTableOutputWriter) WriteCreateOutput(mnemonic string, seed []byte) error {
	tw := table.NewWriter()
	tw.SetOutputMirror(os.Stdout)
	tw.AppendHeader(table.Row{"Mnemonic", "Seed"})
	tw.AppendRow([]interface{}{mnemonic, fmt.Sprintf("%x", seed)})
	tw.Render()
	return nil
}

func (w KeystoreTableOutputWriter) WriteListOutput(accounts []accounts.Account) error {
	tw := table.NewWriter()
	tw.SetOutputMirror(os.Stdout)
	tw.AppendHeader(table.Row{"#", "Address"})
	for i, account := range accounts {
		tw.AppendRow(table.Row{i + 1, account.Address.Hex()})
	}
	tw.Render()
	return nil
}

// KeyStoreJSONOutputWriter writes keystore output in JSON format.
type KeystoreJSONOutputWriter struct{}

func (w KeystoreJSONOutputWriter) WriteCreateOutput(mnemonic string, seed []byte) error {
	seedInfo := map[string]interface{}{
		"mnemonic": mnemonic,
		"seed":     seed,
	}
	jsonOutput, err := json.Marshal(seedInfo)
	if err != nil {
		return err
	}
	fmt.Println(string(jsonOutput))
	return nil
}

func (w KeystoreJSONOutputWriter) WriteListOutput(accounts []accounts.Account) error {
	accountAddresses := make([]string, len(accounts))
	for i, account := range accounts {
		accountAddresses[i] = account.Address.Hex()
	}
	accountInfo := map[string][]string{
		"accounts": accountAddresses,
	}
	jsonOutput, err := json.Marshal(accountInfo)
	if err != nil {
		return err
	}
	fmt.Println(string(jsonOutput))
	return nil
}
