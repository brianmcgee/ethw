package output

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

// SeedOutputWriter is an interface for writing seed information to different output formats.
type SeedOutputWriter interface {
	WriteOutput(mnemonics []string, seeds [][]byte) error
}

type SeedTextOutputWriter struct{}

func (s SeedTextOutputWriter) WriteOutput(mnemonics []string, seeds [][]byte) error {
	// Implement your text output logic here for multiple seeds
	return nil
}

// SeedTableOutputWriter writes seed output in table format.
type SeedTableOutputWriter struct{}

func (s SeedTableOutputWriter) WriteOutput(mnemonics []string, seeds [][]byte) error {
	tw := table.NewWriter()
	tw.SetOutputMirror(os.Stdout)
	tw.AppendHeader(table.Row{"#", "Mnemonic", "Seed"})

	for i := range mnemonics {
		tw.AppendRow([]interface{}{i + 1, mnemonics[i], fmt.Sprintf("%x", seeds[i])})
	}

	tw.Render()
	return nil
}

// SeedJSONOutputWriter writes seed output in JSON format.
type SeedJSONOutputWriter struct{}

func (s SeedJSONOutputWriter) WriteOutput(mnemonics []string, seeds [][]byte) error {
	seedInfo := make([]map[string]interface{}, len(mnemonics))

	for i := range mnemonics {
		seedInfo[i] = map[string]interface{}{
			"mnemonic": mnemonics[i],
			"seed":     seeds[i],
		}
	}

	jsonOutput, err := json.Marshal(seedInfo)
	if err != nil {
		return err
	}

	fmt.Println(string(jsonOutput))
	return nil
}
