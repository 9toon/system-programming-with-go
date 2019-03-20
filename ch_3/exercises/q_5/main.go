package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

func myCopyN(dst io.Writer, src io.Reader, n int64) (written int64, err error) {
	buffer := make([]byte, n)
	size, err := src.Read(buffer)
	if err != nil {
		return 0, err
	}
	fmt.Fprint(dst, string(buffer))

	return int64(size), nil
}

func main() {
	myCopyN(os.Stdout, rand.Reader, 10)
}
