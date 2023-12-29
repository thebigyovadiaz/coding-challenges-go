package fibonacci

import "fmt"

// Fibonacci calculates the nth Fibonacci number using dynamic programming.
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	// Create a slice to store Fibonacci numbers
	fib := make([]int, n+1)
	fib[0], fib[1] = 0, 1

	// Fill the slice using dynamic programming
	for i := 2; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	return fib[n]
}

func ResultFibonacci() {
	// Example usage:
	n := 8
	result := Fibonacci(n)

	fmt.Printf("FIBONACCI >> The %dth Fibonacci number is: %d\n", n, result)
}
