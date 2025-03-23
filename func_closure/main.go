package main

import "fmt"

// Closure untuk menghitung faktorial
func factorial() func(int) int {
	return func(n int) int {
		result := 1
		for i := 1; i <= n; i++ {
			result *= i
		}
		return result
	}
}

func main() {
	fact := factorial()

	for i := 5; i >= 1; i-- {
		fmt.Printf("Faktorial dari %d adalah %d\n", i, fact(i))
	}
}
