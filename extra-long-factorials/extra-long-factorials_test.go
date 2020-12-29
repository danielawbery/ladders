package main

import (
	"fmt"
	"testing"
)

func TestCalcFactorial(t *testing.T) {
	var tests = []struct {
		n    int
		want string
	}{
		// In range cases
		{1, "Result: 1"},
		{3, "Result: 6"},
		{25, "Result: 15511210043330985984000000"},
		{30, "Result: 265252859812191058636308480000000"},
		{100, "Result: 93326215443944152681699238856266700490715968264381621468592963895217599993229915608941463976156518286253697920827223758251185210916864000000000000000000000000"},
		// Error cases
		{-1, "Please enter a valid number. Only 1 <= n <= 100 supported."},
		{0, "Please enter a valid number. Only 1 <= n <= 100 supported."},
		{101, "Please enter a valid number. Only 1 <= n <= 100 supported."},
	}

	for _, ts := range tests {

		testname := fmt.Sprintf("%d", ts.n)
		t.Run(testname, func(t *testing.T) {
			ans := requestFactorial(ts.n)
			if ans != ts.want {
				t.Errorf("got '%v', want '%v'", ans, ts.want)
			}
		})
	}
}
