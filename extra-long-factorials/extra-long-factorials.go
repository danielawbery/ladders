package main

import (
	"fmt"
	"math/big"
)

func requestFactorial(value int) string {
	resultString := "Please enter a valid number. Only 1 <= n <= 100 supported."

	if value >= 1 && (value <= 100) {
		bigValue := big.NewInt(int64(value))
		result := calculateFactorial(bigValue)
		resultString = "Result: " + result.Text(10)
	}

	return resultString
}

func calculateFactorial(bigN *big.Int) (result *big.Int) {
	result = new(big.Int)

	if bigN.Cmp(big.NewInt(0)) == 0 {
		result.SetInt64(1)
		return result
	}
	result.Set(bigN)
	var minusOne big.Int
	minusOne.SetInt64(1)
	result.Mul(result, calculateFactorial(bigN.Sub(bigN, &minusOne)))
	return result
}

func main() {
	fmt.Print(requestFactorial(25))
}
