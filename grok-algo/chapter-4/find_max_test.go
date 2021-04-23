package main

import (
	"testing"
)

func simpleMax(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func findMax(array []int) int {
	if len(array) == 0 {
		return 0
	}

	return simpleMax(array[0], findMax(array[1:]))
}

func TestFindMax(t *testing.T) {
	testCases := []struct {
		desc   string
		input  []int
		output int
	}{
		{
			desc:   "1: Find max value in array (full array)",
			input:  []int{1, 4, 5, 3, 2},
			output: 5,
		},
		{
			desc:   "2: Find max value in array (full array, bigger numbers, unsorted)",
			input:  []int{1, 100, 9, 2, 0, 23, 59},
			output: 100,
		},
		{
			desc:   "3: Find max value in array (full array, bigger numbers, ascending)",
			input:  []int{1, 9, 2, 0, 23, 59, 100},
			output: 100,
		},
		{
			desc:   "4: Find max value in array (single element array)",
			input:  []int{3},
			output: 3,
		},
		{
			desc:   "5: Find max value in array (null array)",
			input:  []int{},
			output: 0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			r := findMax(tC.input)

			if r != tC.output {
				t.Errorf("Expected %d, got %d", tC.output, r)
			}
		})
	}
}
