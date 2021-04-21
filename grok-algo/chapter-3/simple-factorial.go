package main

func factorial(x int) int {
	if x == 1 {
		return 1
	} else {
		return x * factorial(x-1)
	}
}

func main() {
	println(factorial(3))
}
