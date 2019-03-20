package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Create("temp.zip")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	reader := strings.NewReader("content of file in zip\n")
	writer, err := zipWriter.Create("newfile.txt")
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(writer, reader)
	if err != nil {
		panic(err)
	}
}
