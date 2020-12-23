package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_A(t *testing.T) {
	testCases := []struct {
		cups           string
		expectedOutput string
	}{
		{"389125467", "67384529"},
		{"467528193", "43769582"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.cups, func(t *testing.T) {
			cups := load(testCase.cups)
			actualOutput := solveA(cups)
			require.Equal(t, testCase.expectedOutput, actualOutput)
		})
	}
}
