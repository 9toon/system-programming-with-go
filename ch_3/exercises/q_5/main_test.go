package main

import (
	"crypto/rand"
	"io/ioutil"
	"os"
	"testing"
)

func TestMyCopyN(t *testing.T) {
	var length int64 = 8

	name := "tmpfile.txt"
	file, err := os.Create(name)
	if err != nil {
		t.Fatalf("%v\n", err)
	}
	defer file.Close()
	defer os.Remove(name)

	size, err := myCopyN(file, rand.Reader, length)
	if err != nil {
		t.Fatalf("%v\n", err)
	}
	if size != length {
		t.Fatalf("expected size=%d got size=%d", length, size)
	}

	data, err := ioutil.ReadFile(name)
	if err != nil {
		t.Fatalf("%v\n", err)
	}

	if len(data) != int(length) {
		t.Fatalf("Incorrect file size %d\n", len(data))
	}
}
