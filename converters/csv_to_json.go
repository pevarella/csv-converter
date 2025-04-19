package converters

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

func CSVtoJSON(inputFile string, outputFile string) error {
	f, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	if len(records) < 1 {
		return fmt.Errorf("CSV File without HEADER")
	}

	headers := records[0]

	var data []map[string]string

	for _, row := range records[1:] {
		item := make(map[string]string)
		for i, value := range row {
			item[headers[i]] = value
		}
		data = append(data, item)
	}

	out, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer out.Close()

	encoder := json.NewEncoder(out)
	encoder.SetIndent("", " ")
	return encoder.Encode(data)
}
