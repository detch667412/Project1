package extractfile

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractFromZip(t *testing.T) {
	t.Run("Successful extraction", func(t *testing.T) {
		// Create a temporary directory for testing
		tempDir := t.TempDir()
		defer os.RemoveAll(tempDir)

		// Replace with the path to a sample zip file and the correct password
		zipFilePath := "../resources/zip_test.zip"
		password := "1234567890123456789012345678901234567890123456789012345678901234"
		destinationDir := tempDir + "/out1"

		err := ExtractFromZip(zipFilePath, password, destinationDir)
		assert.NoError(t, err, "Unexpected error during extraction")

		// Add assertions for the extracted files if needed
	})
	t.Run("Successful extraction without password", func(t *testing.T) {
		// Create a temporary directory for testing
		tempDir := t.TempDir()
		defer os.RemoveAll(tempDir)

		// Replace with the path to a sample zip file without password
		zipFilePath := "../resources/zip_nopass.zip"
		destinationDir := tempDir + "/out1"

		err := ExtractFromZip(zipFilePath, "", destinationDir)
		assert.NoError(t, err, "Unexpected error during extraction")

		// Add assertions for the extracted files if needed
	})

	t.Run("Invalid zip file", func(t *testing.T) {
		// Replace with the path to an invalid zip file
		invalidZipFilePath := "testdata/invalid.zip"
		password := "testpassword"
		destinationDir := t.TempDir()

		err := ExtractFromZip(invalidZipFilePath, password, destinationDir)
		assert.Error(t, err, "Expected error for invalid zip file")
	})

	t.Run("Incorrect password", func(t *testing.T) {
		// Replace with the path to a sample zip file and an incorrect password
		zipFilePath := "../resources/zip_test.zip"
		incorrectPassword := "wrongpassword"
		destinationDir := t.TempDir()

		err := ExtractFromZip(zipFilePath, incorrectPassword, destinationDir)
		assert.Error(t, err, "Expected error for incorrect password")
	})

	t.Run("Destination directory creation error", func(t *testing.T) {
		// Replace with the path to a sample zip file, correct password, and an invalid destination directory
		zipFilePath := "../resources/zip_test.zip"
		password := "1234567890123456789012345678901234567890123456789012345678901234"
		invalidDestinationDir := "/root/invalid"

		err := ExtractFromZip(zipFilePath, password, invalidDestinationDir)
		assert.Error(t, err, "Expected error for invalid destination directory")
	})
}
