package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func main() {
	source := [][]string{
		{"hello", "world"},
		{"good", "bye"},
	}
	WriteCsv(os.Stdout, source)
}

func WriteCsv(output io.Writer, source [][]string) {
	writer := csv.NewWriter(output)
	for _, line := range source {
		if err := writer.Write(line); err != nil {
			log.Fatalln("error on writing a line to csv: ", err)
		}
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		log.Fatal(err)
	}
}
