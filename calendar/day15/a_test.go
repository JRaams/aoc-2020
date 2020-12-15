package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_A(t *testing.T) {
	testCases := []struct {
		input          string
		expectedOutput int
	}{
		{"1,3,2", 1},
		{"2,1,3", 10},
		{"1,2,3", 27},
		{"2,3,1", 78},
		{"3,2,1", 438},
		{"3,1,2", 1836},
	}

	for _, testCase := range testCases {
		t.Run(testCase.input, func(t *testing.T) {
			input := loadInput(testCase.input)
			result := findLastSpokenNumber(input, 2020)
			require.Equal(t, testCase.expectedOutput, result)
		})
	}
}
