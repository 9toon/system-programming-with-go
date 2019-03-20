package main

import (
	"flag"
	"io"
	"os"
)

func main() {
	srcName := "old.txt"
	dstName := flag.String("dst", "new.txt", "dst filename")
	flag.Parse()

	copy(srcName, *dstName)
}

func copy(srcName, dstName string) {
	src, err := os.Open(srcName)
	if err != nil {
		panic(err)
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	io.Copy(dst, src)
}
