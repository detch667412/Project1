package readfile

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenCSVFile(t *testing.T) {
	t.Run("Open existing file", func(t *testing.T) {
		var errMsg error
		tmpfile, err := os.CreateTemp("", "testfile.csv")
		if err != nil {
			t.Fatal(err)
		}
		defer os.Remove(tmpfile.Name()) // clean up

		file, err := OpenCSVFile(tmpfile.Name())
		errMsg = err
		defer file.Close()

		assert.NoError(t, errMsg, "Unexpected error when opening an existing file")
		assert.NotNil(t, file, "File handle is nil")
	})

	t.Run("Open non-existing file", func(t *testing.T) {
		// Provide a non-existing file path
		nonExistingFile := "nonexistent.csv"

		file, err := OpenCSVFile(nonExistingFile)

		// No need to defer file.Close() here, as file is expected to be nil
		assert.Error(t, err, "Expected an error when opening a non-existing file")
		assert.Nil(t, file, "File handle is not nil for a non-existing file")
	})
}

func TestCreateCSVReader(t *testing.T) {
	t.Run("Create CSV reader with delimiter", func(t *testing.T) {
		tmpfile, err := os.CreateTemp("", "testfile.csv")
		if err != nil {
			t.Fatal(err)
		}
		defer os.Remove(tmpfile.Name()) // clean up

		file, err := os.Open(tmpfile.Name())
		if err != nil {
			t.Fatal(err)
		}
		defer file.Close()

		reader := CreateCSVReader(file, '|')

		assert.Equal(t, '|', reader.Comma, "Unexpected delimiter in CSV reader")
	})
}
func TestReadCSVHeaderFromFile(t *testing.T) {
	t.Run("Read empty CSV", func(t *testing.T) {
		tmpfile, err := os.CreateTemp("", "testfile.csv")
		if err != nil {
			t.Fatal(err)
		}
		defer os.Remove(tmpfile.Name()) // clean up

		file, err := os.Open(tmpfile.Name())
		if err != nil {
			t.Fatal(err)
		}
		defer file.Close()

		reader := CreateCSVReader(file, '|')

		records, err := ReadRecords(reader)
		assert.NoError(t, err, "Unexpected error when reading all records from an empty CSV")
		assert.Empty(t, records, "Expected an empty slice when reading from an empty CSV")
	})
	t.Run("Check header of input file(csv)", func(t *testing.T) {
		filename := "../input.csv"
		file, err := os.Open(filename)
		if err != nil {
			t.Fatalf("Error opening file %s: %v", filename, err)
		}
		defer file.Close()

		reader := CreateCSVReader(file, '|')

		// Use Read to read just the header
		actualHeader, err := reader.Read()
		if err != nil {
			t.Fatalf("Error reading header: %v", err)
		}

		expectedHeader := []string{"Id", "Name", "Address", "Tel"}

		assert.Equal(t, expectedHeader, actualHeader, "Mismatch in CSV header")
	})
}
