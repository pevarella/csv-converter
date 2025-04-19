package main

import (
	"fmt"
	"os"

	"github.com/pevarella/csv-converter/converters"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("csv-converter <csv_file> <json|parquet> <output_file>")
		return
	}

	inputFile := os.Args[1]
	// format := os.Args[2]
	outputFile := os.Args[3]

	err := converters.CSVtoJSON(inputFile, outputFile)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Sucess!")
	}
}
