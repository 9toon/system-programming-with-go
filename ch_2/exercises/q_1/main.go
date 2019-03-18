package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	WriteFile("tempfile.txt")
}

func WriteFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Fprintf(file, "%04d, %s, %2.2f\n", 10, "test", 0.2)
	file.Close()
}
