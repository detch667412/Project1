package encryptzip

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryptZip(t *testing.T) {
	t.Run("Encrypt and zip an existing file", func(t *testing.T) {
		// Create a temporary test file
		testFileName := "testfile.txt"
		err := createTestFile(testFileName)
		assert.NoError(t, err, "Error creating test file")
		defer deleteTestFile(testFileName)
		defer os.Remove("output.zip") //delete zip's test file

		// Test the function
		err = EncryptZip(testFileName)
		assert.NoError(t, err, "Error encrypting and zipping the file")
	})

	t.Run("Fail to open non-existing file", func(t *testing.T) {
		// Test case 2: File not found
		err := EncryptZip("nonexistentfile.txt")
		assert.Error(t, err, "Expected error for nonexistent file")
	})
}

// Helper function to create a test file with some content
func createTestFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("Test content for encryption")
	return err
}

// Helper function to delete the test file
func deleteTestFile(filename string) {
	_ = os.Remove(filename)
}
