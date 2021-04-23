package main

import "testing"

func count_elements(array []int) int {
	if len(array) == 0 {
		return 0
	}

	return 1 + count_elements(array[1:])
}

func Test_Count_Elements(t *testing.T) {
	testCases := []struct {
		desc   string
		input  []int
		output int
	}{
		{
			desc:   "1: Count array elements (full array)",
			input:  []int{1, 2, 3, 4, 5},
			output: 5,
		},
		{
			desc:   "2: Count array elements (single element array)",
			input:  []int{1},
			output: 1,
		},
		{
			desc:   "3: Count array elements (null array)",
			input:  []int{},
			output: 0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			r := count_elements(tC.input)

			if r != tC.output {
				t.Errorf("Expected %d, got %d", tC.output, r)
			}
		})
	}
}
