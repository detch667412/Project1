package main

import (
	"io"
	"log"
	"os"

	"github.com/yeka/zip"
)

func main() {
	err := EncryptZip("input.csv")
	if err != nil {
		log.Fatal(err)
	}

	//Dont forget to write unit tesr for "EncryptZip"
}

func EncryptZip(fileToZip string) error {
	// Open the file to be zippedd
	file, err := os.Open(fileToZip)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a new ZIP file
	zipFile, err := os.Create("output.zip")
	if err != nil {
		return err
	}
	defer zipFile.Close()

	// Create a ZIP writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	w, err := zipWriter.Encrypt(fileToZip, `golang`, zip.AES256Encryption)
	if err != nil {
		return err
	}

	// Copy the content of the existing file to the ZIP archive
	_, err = io.Copy(w, file)
	if err != nil {
		return err
	}

	log.Printf("File '%s' successfully zipped to 'output.zip'\n", fileToZip)
	return nil
}
