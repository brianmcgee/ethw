package output

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

// SeedOutputWriter is an interface for writing seed information to different output formats.
type SeedOutputWriter interface {
	WriteOutput(mnemonics []string, seeds [][]byte) error
}

// SeedTextOutputWriter writes seed output in text format.
type SeedTextOutputWriter struct{}

func (s SeedTextOutputWriter) WriteOutput(mnemonics []string, seeds [][]byte) error {
	if len(mnemonics) == 0 {
		fmt.Println("No mnemonics found.")
		return nil
	}

	fmt.Println("Seed Information:")
	for i, mnemonic := range mnemonics {
		fmt.Printf("  Entry #%d\n", i+1)
		fmt.Printf("  Seed: %x\n\n", seeds[i])
		fmt.Printf("  Mnemonic: %s\n", mnemonic)
	}

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

// SeedCSVOutputWriter writes seed output in CSV format.
type SeedCSVOutputWriter struct{}

func (s SeedCSVOutputWriter) WriteOutput(mnemonics []string, seeds [][]byte) error {
	csvWriter := csv.NewWriter(os.Stdout)
	defer csvWriter.Flush()

	// Write the CSV header
	if err := csvWriter.Write([]string{"#", "Mnemonic", "Seed"}); err != nil {
		return fmt.Errorf("writing CSV header: %w", err)
	}

	for i, mnemonic := range mnemonics {
		seedHex := fmt.Sprintf("%x", seeds[i])
		record := []string{fmt.Sprintf("%d", i+1), mnemonic, seedHex}

		// Write each record
		if err := csvWriter.Write(record); err != nil {
			return fmt.Errorf("writing CSV record: %w", err)
		}
	}

	return nil
}
