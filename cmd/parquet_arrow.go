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
	if err := parquetArrowCmd.MarkFlagRequired("input"); err != nil {
		fmt.Printf("Error setting required flag: %v\n", err)
	}
	if err := parquetArrowCmd.MarkFlagRequired("output"); err != nil {
		fmt.Printf("Error setting required flag: %v\n", err)
	}
	rootCmd.AddCommand(parquetArrowCmd)
}
