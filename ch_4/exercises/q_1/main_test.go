package main

import "testing"

func TestNotify(t *testing.T) {
	var tests = []struct {
		s        int
		expected string
	}{
		{1, "Passed 1 sec"},
		{0, "Passed 0 sec"},
	}

	for _, test := range tests {
		got := notify(test.s)
		if got != test.expected {
			t.Errorf("result = %q\nexpected = %q", got, test.expected)
		}
	}
}
