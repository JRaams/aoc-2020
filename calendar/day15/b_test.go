package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_B(t *testing.T) {
	testCases := []struct {
		input          string
		expectedOutput int
	}{
		{"0,3,6", 175594},
		{"1,3,2", 2578},
		{"2,1,3", 3544142},
		{"1,2,3", 261214},
		{"2,3,1", 6895259},
		{"3,2,1", 18},
		{"3,1,2", 362},
	}

	for _, testCase := range testCases {
		t.Run(testCase.input, func(t *testing.T) {
			t.Parallel()

			input := loadInput(testCase.input)
			result := findLastSpokenNumber(input, 30000000)
			require.Equal(t, testCase.expectedOutput, result)
		})
	}
}
