package main

import (
	"reflect"
	"testing"
)

func quickSort(array []int) []int {
	if len(array) < 2 {
		return array
	}

	pivot := array[0]
	// DON'T FORGET TO REMOVE PIVOT FROM THE ARRAY!!!!
	array = array[1:]
	lesser := []int{}
	greater := []int{}

	for _, i := range array {
		if i <= pivot {
			lesser = append(lesser, i)
		} else {
			greater = append(greater, i)
		}
	}

	lesser = quickSort(lesser)
	lesser = append(lesser, pivot)
	greater = quickSort(greater)

	return append(lesser, greater...)
}

func TestQuickSort(t *testing.T) {
	testCases := []struct {
		desc   string
		input  []int
		output []int
	}{
		{
			desc:   "1: QuickSort (> 2 element array)",
			input:  []int{3, 5, 2, 1, 4},
			output: []int{1, 2, 3, 4, 5},
		},
		{
			desc:   "2: QuickSort (2 element array)",
			input:  []int{7, 3},
			output: []int{3, 7},
		},
		{
			desc:   "3: QuickSort (single element array)",
			input:  []int{1},
			output: []int{1},
		},
		{
			desc:   "4: QuickSort (null array)",
			input:  []int{},
			output: []int{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			r := quickSort(tC.input)

			if reflect.DeepEqual(r, tC.output) != true {
				t.Errorf("Expected %d, got %d", tC.output, r)
			}
		})
	}

}
