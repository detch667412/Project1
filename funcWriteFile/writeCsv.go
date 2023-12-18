package writefile

import (
	"encoding/csv"
	"os"
	"reflect"
	"strconv"
)

type Person struct {
	Name string `csv:"Name"`
	Age  int    `csv:"Age"`
	City string `csv:"City"`
}

func WriteStructsToCSV(filename string, people []Person) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Set the delimiter to '|'
	writer.Comma = '|'

	// Write header
	header := getCSVHeader(people)
	if err := writeCSVRow(writer, header); err != nil {
		return err
	}

	// Write data
	for _, person := range people {
		row := structToCSVRow(person)
		if err := writeCSVRow(writer, row); err != nil {
			return err
		}
	}

	return nil
}

func getCSVHeader(people []Person) []string {
	if len(people) == 0 {
		return nil
	}

	header := []string{}
	t := reflect.TypeOf(people[0])
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("csv")
		header = append(header, tag)
	}

	return header
}

func structToCSVRow(person Person) []string {
	return []string{person.Name, strconv.Itoa(person.Age), person.City}
}

func writeCSVRow(writer *csv.Writer, row []string) error {
	if err := writer.Write(row); err != nil {
		return err
	}
	return nil
}
