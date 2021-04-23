package main

import "testing"

func loop_sum(array []int) int {
	total := 0
	for i := 1; i < len(array)+1; i++ {
		total += i
	}
	return total
}

func recursion_sum(array []int) int {
	if len(array) == 0 {
		return 0
	}

	return array[0] + recursion_sum(array[1:])
}

func Test_Sum_Array(t *testing.T) {
	full_array := []int{1, 2, 3, 4, 5}
	full_array_output := 15

	single_element_array := []int{1}
	single_element_array_output := 1

	null_array := []int{}
	null_array_output := 0

	testCasesLoop := []struct {
		desc   string
		input  []int
		output int
	}{
		{
			desc:   "1: Sum array by Loop (full array)",
			input:  full_array,
			output: full_array_output,
		},
		{
			desc:   "2: Sum array by Loop (single element array)",
			input:  single_element_array,
			output: single_element_array_output,
		},
		{
			desc:   "3: Sum array by Loop (null array)",
			input:  null_array,
			output: null_array_output,
		},
	}
	testCasesRecursion := []struct {
		desc   string
		input  []int
		output int
	}{
		{
			desc:   "1: Sum array by Recursion (full array)",
			input:  full_array,
			output: full_array_output,
		},
		{
			desc:   "2: Sum array by Recursion (single element array)",
			input:  single_element_array,
			output: single_element_array_output,
		},
		{
			desc:   "3: Sum array by Recursion (null array)",
			input:  null_array,
			output: null_array_output,
		},
	}
	for _, tC := range testCasesLoop {
		t.Run(tC.desc, func(t *testing.T) {
			r := loop_sum(tC.input)

			if r != tC.output {
				t.Errorf("Expected %d, got %d", tC.output, r)
			}
		})
	}
	for _, tC := range testCasesRecursion {
		t.Run(tC.desc, func(t *testing.T) {
			r := recursion_sum(tC.input)

			if r != tC.output {
				t.Errorf("Expected %d, got %d", tC.output, r)
			}
		})
	}
}
