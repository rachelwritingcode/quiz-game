package csv

import (
	"encoding/csv"
	"flag"
	"os"
)

func SetCSV() string {

	csvFilePath := flag.String("csv", "problems.csv", "filepath for problem set csv file")
	flag.Parse()
	return *csvFilePath
}

func ReadCSV(csvFilePath string) []string {

	csvRecords := []string{}
	f, err := os.Open(csvFilePath)
	if err != nil {
		csvRecords = append(csvRecords, "error")
		return csvRecords
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	fileRecords, err := csvReader.ReadAll()
	if err != nil {
		csvRecords = append(csvRecords, "error")
		return csvRecords
	}
	for _, s := range fileRecords {
		csvRecords = append(csvRecords, s...)
	}
	return csvRecords
}
