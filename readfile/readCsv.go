package readfile

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

// OpenCSVFile opens a CSV file and returns a file handle.
func OpenCSVFile(filename string) (*os.File, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	return file, nil
}

// CreateCSVReader creates a new CSV reader for a given file handle with a specified delimiter.
func CreateCSVReader(file *os.File, delimiter rune) *csv.Reader {
	reader := csv.NewReader(file)
	reader.Comma = delimiter
	return reader
}

// ReadAllRecords reads all records from a CSV reader.
func ReadRecords(reader *csv.Reader) ([][]string, error) {
	var records [][]string

	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break // End of file
			}
			return nil, fmt.Errorf("error reading CSV: %v", err)
		}
		records = append(records, record)
	}

	return records, nil
}

// ReadCSVFile reads all records from a CSV file with a specified delimiter.
func ReadCSVFile(filename string, delimiter rune) ([][]string, error) {
	file, err := OpenCSVFile(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := CreateCSVReader(file, delimiter)
	records, err := ReadRecords(reader)
	if err != nil {
		return nil, err
	}

	return records, nil
}
