package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "csv-converter",
	Short: "Tool for convert CSV into JSON/Parquet",
	Long:  "CSV-Converter reads the csv file and generate a JSON or a Parquet file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use sub: json or parquet")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
