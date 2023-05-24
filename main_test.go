package main

import (
	"testing"
)

func TestSplit(t *testing.T) {
	testCases := []struct {
		input    []any
		expected [][]byte
	}{
		{
			input: []any{
				[]byte("Hello\nWorld"),
				byte('\n'),
			},
			expected: [][]byte{
				[]byte("Hello"),
				[]byte("World"),
			},
		},
	}

	for _, tc := range testCases {
		actual := Split(tc.input[0].([]byte), tc.input[1].(byte))

		expected := string(tc.expected[0])
		got := string(actual[0])

		if expected != got {
			t.Errorf("Expected %s, got %s", expected, got)
		}
	}
}
