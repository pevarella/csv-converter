package cmd

import (
	"fmt"

	"github.com/pevarella/csv-converter/converters"
	"github.com/spf13/cobra"
)

var (
	inFile, outFile string
)

var jsonCmd = &cobra.Command{
	Use:   "json",
	Short: "Convert CSV to JSON",
	RunE: func(cmd *cobra.Command, args []string) error {
		if inFile == "" || outFile == "" {
			return fmt.Errorf("flags --input and --output are required")
		}
		return converters.CSVtoJSON(inFile, outFile)
	},
}

func init() {
	jsonCmd.Flags().StringVarP(&inFile, "input", "i", "", "Path to CSV file")
	jsonCmd.Flags().StringVarP(&outFile, "output", "o", "", "Path to JSON file")
	jsonCmd.MarkFlagRequired("input")
	jsonCmd.MarkFlagRequired("output")
	rootCmd.AddCommand(jsonCmd)
}
