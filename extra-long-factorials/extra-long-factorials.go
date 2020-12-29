package main

import (
	"fmt"
	"math/big"
)

// requestFactorial takes an integer value (within the range 1 <= n <= 100 and returns either:
// 1) A result string with the factorial of that integer
// 2) (When outside the range) an error message requesting a valid number.

func requestFactorial(value int) string {

	// Default to failure string and only change value if requested numnber falls in bounds.
	resultString := "Please enter a valid number. Only 1 <= n <= 100 supported."

	// Check if number is in bounds.
	if value >= 1 && (value <= 100) {

		// Create a new big.Int from value. We have to case the int to int64 to call big.NewInt().
		bigValue := big.NewInt(int64(value))
		result := calculateFactorial(bigValue)
		resultString = "Result: " + result.Text(10)
	}

	return resultString
}

// calculateFactorial takes a requested number as *big.Int and returns a result as *big.Int that is its factorial.
// calculateFactorial is recursive and calls itself until request = 0.
func calculateFactorial(request *big.Int) (result *big.Int) {

	// Create new big.Int to hold the result value.
	result = new(big.Int)

	// Base case: Check whether the result is equal to 0, using Cmp() and return 1 if so.
	if request.Cmp(big.NewInt(0)) == 0 {
		result.SetInt64(1)
		return result
	}

	// Set result to value of request. Create decrement as *big.Int to calculate request - 1 for factorial.
	result.Set(request)
	decrement := big.NewInt(1)

	// Use Mul() from big to multiply current value of result with the outcome of a recursive call, with result - 1.
	result.Mul(result, calculateFactorial(request.Sub(request, decrement)))
	return result
}

func main() {
	fmt.Print(requestFactorial(25))
}
