package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	// `Example of...` から `Section` のみを抜き出す処理
	reader := strings.NewReader("Example of io.SectionReader\n")
	sectionReader := io.NewSectionReader(reader, 14, 7)
	io.Copy(os.Stdout, sectionReader)
}
