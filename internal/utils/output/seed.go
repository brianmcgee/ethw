package output

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

// SeedOutputWriter is an interface for writing seed information to different output formats.
type SeedOutputWriter interface {
	WriteOutput(mnemonic string, seed []byte) error
}

// SeedTableDataWriter writes seed output in table format.
type SeedTableDataWriter struct{}

func (s SeedTableDataWriter) WriteOutput(mnemonic string, seed []byte) error {
	tw := table.NewWriter()
	tw.SetOutputMirror(os.Stdout)
	tw.AppendHeader(table.Row{"Mnemonic", "Seed"})
	tw.AppendRow([]interface{}{mnemonic, fmt.Sprintf("%x", seed)})
	tw.Render()
	return nil
}

// SeedJSONDataWriter writes seed output in JSON format.
type SeedJSONDataWriter struct{}

func (s SeedJSONDataWriter) WriteOutput(mnemonic string, seed []byte) error {
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
