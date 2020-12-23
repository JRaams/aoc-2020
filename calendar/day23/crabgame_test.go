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
			vc, currentCup := load(testCase.cups, false)
			actualOutput := solveA(vc, currentCup)
			require.Equal(t, testCase.expectedOutput, actualOutput)
		})
	}
}

func Test_B(t *testing.T) {
	testCases := []struct {
		cups           string
		expectedOutput int
	}{
		{"389125467", 149245887792},
	}

	for _, testCase := range testCases {
		t.Run(testCase.cups, func(t *testing.T) {
			vc, currentCup := load(testCase.cups, true)
			actualOutput := solveB(vc, currentCup)
			require.Equal(t, testCase.expectedOutput, actualOutput)
		})
	}
}
