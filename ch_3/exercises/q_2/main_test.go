package main

import (
	"io/ioutil"
	"testing"
)

func TestCreateTmpFile(t *testing.T) {
	name := "tmpfile.txt"
	err := createTmpFile(name)
	if err != nil {
		t.Fatalf("%v\n", err)
	}
	data, err := ioutil.ReadFile(name)
	if err != nil {
		t.Fatalf("%v\n", err)
	}

	if len(data) != 1024 {
		t.Fatalf("Incorrect file size %d\n", len(data))
	}
}
