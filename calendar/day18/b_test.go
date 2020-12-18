package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_B(t *testing.T) {
	testCases := []struct {
		line           string
		expectedOutput int
	}{
		{"1 + 2 * 3 + 4 * 5 + 6", 231},
		{"1 + (2 * 3) + (4 * (5 + 6))", 51},
		{"2 * 3 + (4 * 5)", 46},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 1445},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 669060},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 23340},
	}

	for _, testCase := range testCases {
		t.Run(testCase.line, func(t *testing.T) {
			actualOutput := evaluate(testCase.line, true)
			require.Equal(t, testCase.expectedOutput, actualOutput)
		})
	}
}
