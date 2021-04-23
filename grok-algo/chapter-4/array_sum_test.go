package main

import "testing"

func loopSum(array []int) int {
	total := 0
	for i := 1; i < len(array)+1; i++ {
		total += i
	}
	return total
}

func recursionSum(array []int) int {
	if len(array) == 0 {
		return 0
	}

	return array[0] + recursionSum(array[1:])
}

func TestSumArray(t *testing.T) {
	fullArray := []int{1, 2, 3, 4, 5}
	fullArrayOutput := 15

	singleElementArray := []int{1}
	singleElementArrayOutput := 1

	nullArray := []int{}
	nullArrayOutput := 0

	testCasesLoop := []struct {
		desc   string
		input  []int
		output int
	}{
		{
			desc:   "1: Sum array by Loop (full array)",
			input:  fullArray,
			output: fullArrayOutput,
		},
		{
			desc:   "2: Sum array by Loop (single element array)",
			input:  singleElementArray,
			output: singleElementArrayOutput,
		},
		{
			desc:   "3: Sum array by Loop (null array)",
			input:  nullArray,
			output: nullArrayOutput,
		},
	}
	testCasesRecursion := []struct {
		desc   string
		input  []int
		output int
	}{
		{
			desc:   "1: Sum array by Recursion (full array)",
			input:  fullArray,
			output: fullArrayOutput,
		},
		{
			desc:   "2: Sum array by Recursion (single element array)",
			input:  singleElementArray,
			output: singleElementArrayOutput,
		},
		{
			desc:   "3: Sum array by Recursion (null array)",
			input:  nullArray,
			output: nullArrayOutput,
		},
	}
	for _, tC := range testCasesLoop {
		t.Run(tC.desc, func(t *testing.T) {
			r := loopSum(tC.input)

			if r != tC.output {
				t.Errorf("Expected %d, got %d", tC.output, r)
			}
		})
	}
	for _, tC := range testCasesRecursion {
		t.Run(tC.desc, func(t *testing.T) {
			r := recursionSum(tC.input)

			if r != tC.output {
				t.Errorf("Expected %d, got %d", tC.output, r)
			}
		})
	}
}
