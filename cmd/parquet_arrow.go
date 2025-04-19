package cmd

import (
	"fmt"

	"github.com/pevarella/csv-converter/converters"
	"github.com/spf13/cobra"
)

var parquetArrowCmd = &cobra.Command{
	Use:   "parquet-arrow",
	Short: "Convert CSV to Parquet",
	RunE: func(cmd *cobra.Command, args []string) error {
		if inFile == "" || outFile == "" {
			return fmt.Errorf("flags --input and --output are required")
		}
		return converters.CSVtoParquetArrow(inFile, outFile)
	},
}

func init() {
	parquetArrowCmd.Flags().StringVarP(&inFile, "input", "i", "", "Path to CSV file")
	parquetArrowCmd.Flags().StringVarP(&outFile, "output", "o", "", "Path to Parquet file")
	parquetArrowCmd.MarkFlagRequired("input")
	parquetArrowCmd.MarkFlagRequired("output")
	rootCmd.AddCommand(parquetArrowCmd)
}
