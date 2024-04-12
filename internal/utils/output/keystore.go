package output

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aldoborrero/ethw/internal/keystore"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/jedib0t/go-pretty/v6/table"
)

// KeystoreOutputWriter is an interface for writing keystore information to different output formats.
type KeystoreOutputWriter interface {
	WriteCreateOutput(ks keystore.KeystoreWrapper) error
	WriteListOutput(accounts []accounts.Account) error
}

// KeystoreTextOutputWriter writes keystore output in pure text format.
type KeystoreTextOutputWriter struct{}

func (w KeystoreTextOutputWriter) WriteCreateOutput(ks keystore.KeystoreWrapper) error {
	fmt.Println("Account Creation Details:")
	for _, account := range ks.Accounts() {
		fmt.Printf("Address: %s\nKeystore Path: %s\n\n", account.Address.Hex(), account.URL.Path)
	}
	return nil
}

func (w KeystoreTextOutputWriter) WriteListOutput(accounts []accounts.Account) error {
	if len(accounts) == 0 {
		fmt.Println("No accounts found.")
		return nil
	}

	fmt.Println("List of Wallets:")
	for i, account := range accounts {
		fmt.Printf("Wallet %d: %s\n", i+1, account.Address.Hex())
	}

	return nil
}

// KeystoreTableOutputWriter writes keystore output in table format.
type KeystoreTableOutputWriter struct{}

func (w KeystoreTableOutputWriter) WriteCreateOutput(ks keystore.KeystoreWrapper) error {
	tw := table.NewWriter()
	tw.SetOutputMirror(os.Stdout)
	tw.AppendHeader(table.Row{"#", "Address", "Keystore Path"})
	for i, account := range ks.Accounts() {
		tw.AppendRow(table.Row{i + 1, account.Address.Hex(), account.URL.Path})
	}
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

func (w KeystoreJSONOutputWriter) WriteCreateOutput(ks keystore.KeystoreWrapper) error {
	accounts := ks.Accounts()
	keystoreInfo := make([]map[string]string, len(accounts))
	for i, account := range accounts {
		keystoreInfo[i] = map[string]string{
			"address":       account.Address.Hex(),
			"keystore_path": account.URL.Path,
		}
	}
	jsonOutput, err := json.Marshal(keystoreInfo)
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

// KeystoreCSVOutputWriter writes keystore output in CSV format.
type KeystoreCSVOutputWriter struct{}

func (w KeystoreCSVOutputWriter) WriteCreateOutput(ks keystore.KeystoreWrapper) error {
	csvWriter := csv.NewWriter(os.Stdout)
	defer csvWriter.Flush()

	err := csvWriter.Write([]string{"Address", "Keystore Path"})
	if err != nil {
		return err
	}

	for _, account := range ks.Accounts() {
		err := csvWriter.Write([]string{account.Address.Hex(), account.URL.Path})
		if err != nil {
			return err
		}
	}

	return nil
}

func (w KeystoreCSVOutputWriter) WriteListOutput(accounts []accounts.Account) error {
	csvWriter := csv.NewWriter(os.Stdout)
	defer csvWriter.Flush()

	err := csvWriter.Write([]string{"Index", "Address"})
	if err != nil {
		return err
	}

	for i, account := range accounts {
		err := csvWriter.Write([]string{fmt.Sprintf("%d", i+1), account.Address.Hex()})
		if err != nil {
			return err
		}
	}

	return nil
}
