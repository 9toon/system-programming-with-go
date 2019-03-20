package main

import (
	"crypto/rand"
	"io"
	"os"
)

func main() {
	err := createTmpFile("tmpfile.txt")
	if err != nil {
		panic(err)
	}
}

func createTmpFile(name string) error {
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.CopyN(file, rand.Reader, 1024)
	if err != nil {
		return err
	}

	return nil
}
