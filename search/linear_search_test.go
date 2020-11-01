package search

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_LinearSearch(t *testing.T) {
	testCases := []struct {
		input          []int
		target         int
		expectedOutput bool
	}{
		{
			input:          []int{1, 2, 3, 4},
			target:         2,
			expectedOutput: true,
		},
		{
			input: []int{},
			target: 2,
			expectedOutput: false,
		},
		{
			input: []int{1},
			target: 2,
			expectedOutput: false,
		},
		{
			input: nil,
			target: 2,
			expectedOutput: false,
		},
	}

	for _, tc := range(testCases) {
		out := LinearSearch(tc.input, tc.target)
		assert.Equal(t, tc.expectedOutput, out)
	}
}
