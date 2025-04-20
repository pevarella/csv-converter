package main

import (
	"github.com/pevarella/csv-converter/cmd"
	"github.com/pevarella/csv-converter/logger"
)

func main() {
	logger.Init(true)
	cmd.Execute()
}
