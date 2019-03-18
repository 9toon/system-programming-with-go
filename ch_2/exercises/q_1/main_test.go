package main

import (
	"io/ioutil"
	"testing"
)

func TestWriteFile(t *testing.T) {
	filename := "tempfile.txt"
	expected := "0010, test, 0.20\n"
	WriteFile(filename)
	actual, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("%v", err)
	}
	if string(actual) != expected {
		t.Fatalf("result = %q, expected = %q", actual, expected)
	}
}
