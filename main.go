package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/yeka/zip"
)

func main() {

	encryptedFile := "resources/zip_test.zip"
	password := "1234567890123456789012345678901234567890123456789012345678901234"
	destinationDir := "extracted"

	err := ExtractFromZip(encryptedFile, password, destinationDir)
	if err != nil {
		log.Fatal(err)
	}

}

func ExtractFromZip(encryptedFile, password, destinationDir string) error {
	r, err := zip.OpenReader(encryptedFile)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		if f.IsEncrypted() {
			f.SetPassword(password)
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		// Create the file or directory in the destination directory
		destPath := filepath.Join(destinationDir, f.Name)
		if f.FileInfo().IsDir() {
			_ = os.MkdirAll(destPath, os.ModePerm)
		} else {
			err := os.MkdirAll(filepath.Dir(destPath), os.ModePerm)
			if err != nil {
				return err
			}

			outFile, err := os.Create(destPath)
			if err != nil {
				return err
			}
			defer outFile.Close()

			// Copy the contents of the file to the destination file
			_, err = io.Copy(outFile, rc)
			if err != nil {
				return err
			}

			fmt.Printf("Extracted %v to %v\n", f.Name, destPath)
		}
	}

	return nil
}
