package writefile

import (
	"encoding/csv"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCSVHeader(t *testing.T) {
	t.Run("Get CSV header from Person struct", func(t *testing.T) {
		people := []Person{
			{Name: "John", Age: 30, City: "New York"},
			{Name: "Alice", Age: 25, City: "San Francisco"},
		}

		header := getCSVHeader(people)

		expectedHeader := []string{"Name", "Age", "City"}
		assert.Equal(t, expectedHeader, header, "Unexpected CSV header")
	})

	t.Run("Get CSV header from empty Person slice", func(t *testing.T) {
		people := []Person{} // empty input

		header := getCSVHeader(people)

		expectedHeader := []string(nil) // empty header for empty input
		assert.Equal(t, expectedHeader, header, "Unexpected CSV header for empty input")
	})
}

func TestStructToCSVRow(t *testing.T) {
	t.Run("Convert Person struct to CSV row", func(t *testing.T) {
		person := Person{Name: "John", Age: 30, City: "New York"}

		row := structToCSVRow(person)

		expectedRow := []string{"John", "30", "New York"}
		assert.Equal(t, expectedRow, row, "Unexpected CSV row")
	})
}

func TestWriteCSVRow(t *testing.T) {
	t.Run("Write CSV row to CSV writer", func(t *testing.T) {
		// Create a temporary file
		tmpfile, err := os.CreateTemp("", "testfile.csv")
		if err != nil {
			t.Fatal(err)
		}
		defer os.Remove(tmpfile.Name()) // clean up

		file, err := os.OpenFile(tmpfile.Name(), os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			t.Fatal(err)
		}
		defer file.Close()

		writer := csv.NewWriter(file)

		row := []string{"John", "30", "New York"}

		// Call the function
		err = writeCSVRow(writer, row)
		writer.Flush()
		assert.NoError(t, err, "Error writing CSV row to CSV writer")

		// Ensure the content is immediately flushed to the file
		_ = file.Sync()

		// Check if there's any error reported by the CSV writer
		assert.Nil(t, writer.Error(), "Unexpected error in CSV writer")

		// Read the content of the written file and check if it's correct
		fileContent, err := os.ReadFile(tmpfile.Name())
		assert.NoError(t, err, "Error reading CSV file")

		expectedContent := "John,30,New York\n"
		assert.Equal(t, expectedContent, string(fileContent), "Unexpected content in CSV file")
	})
}

func TestWriteStructsToCSV(t *testing.T) {
	t.Run("Write structs to CSV with custom delimiter", func(t *testing.T) {
		// Create a temporary file
		tmpfile, err := os.CreateTemp("", "testfile.csv")
		if err != nil {
			t.Fatal(err)
		}
		defer os.Remove(tmpfile.Name()) // clean up

		people := []Person{
			{Name: "John", Age: 30, City: "New York"},
			{Name: "Alice", Age: 25, City: "San Francisco"},
		}

		// Call the function
		err = WriteStructsToCSV(tmpfile.Name(), people)
		assert.NoError(t, err, "Error writing structs to CSV")

		// Read the content of the written file and check if it's correct
		fileContent, err := os.ReadFile(tmpfile.Name())
		assert.NoError(t, err, "Error reading CSV file")

		expectedContent := "Name|Age|City\nJohn|30|New York\nAlice|25|San Francisco\n"
		assert.Equal(t, expectedContent, string(fileContent), "Unexpected content in CSV file")
	})

	t.Run("Write structs to CSV with empty input", func(t *testing.T) {
		// Create a temporary file
		tmpfile, err := os.CreateTemp("", "testfile.csv")
		if err != nil {
			t.Fatal(err)
		}
		defer os.Remove(tmpfile.Name()) // clean up

		people := []Person{} // empty input

		// Call the function
		err = WriteStructsToCSV(tmpfile.Name(), people)
		assert.NoError(t, err, "Error writing structs to CSV with empty input")

		// Read the content of the written file and check if it's correct
		fileContent, err := os.ReadFile(tmpfile.Name())
		assert.NoError(t, err, "Error reading CSV file")

		expectedContent := "\n" // empty content for empty input
		assert.Equal(t, expectedContent, string(fileContent), "Unexpected content in CSV file")
	})
}
