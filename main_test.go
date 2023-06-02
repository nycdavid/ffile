package main

import (
	"fmt"
	"testing"
)

func TestSplit(t *testing.T) {
	testCases := []struct {
		name     string
		input    []any
		expected [][]byte
	}{
		{
			name: "happy path",
			input: []any{
				[]byte("Hello-World"),
				byte('-'),
			},
			expected: [][]byte{
				[]byte("Hello"),
				[]byte("World"),
			},
		},
		{
			name: "edge case: ends with delimiter",
			input: []any{
				[]byte("Hello-"),
				byte('-'),
			},
			expected: [][]byte{
				[]byte("Hello"),
			},
		},
		{
			name: "edge case: empty input slice",
			input: []any{
				[]byte{},
				byte('-'),
			},
			expected: [][]byte{},
		},
	}

	for i, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Split(tc.input[0].([]byte), tc.input[1].(byte))

			for j, expectedLine := range tc.expected {
				got := string(actual[j])
				if string(expectedLine) != got {
					t.Errorf("[%d] Expected %s, got %s", i+1, expectedLine, got)
				}
			}

			if len(tc.expected) != len(actual) {
				t.Errorf("[%d] Expected length %d, got %d", i+1, len(tc.expected), len(actual))
			}
		})
	}
}

func TestFindIn(t *testing.T) {
	type input struct {
		fname    string
		term     string
		contents [][]byte
	}

	testCases := []struct {
		name     string
		input    input
		expected string
	}{
		{
			name: "happy path",
			input: input{
				fname: "anon.txt",
				term:  "world",
				contents: [][]byte{
					[]byte("hello"),
					[]byte("world"),
				},
			},
			expected: "anon.txt:2",
		},
	}

	for i, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			chnl := make(chan hit, 1)
			var hits []string
			go FindIn(tc.input.fname, tc.input.term, tc.input.contents, chnl)

			for {
				select {
				case v := <-chnl:
					hits = append(hits, fmt.Sprintf("%s:%d", v.fname, v.line))
				default:
					return
				}
			}

			actual := hits[0]

			if actual != tc.expected {
				t.Errorf("[%d] Expected %s, got %s", i+1, tc.expected, actual)
			}
		})
	}
}
