package main

import "fmt"

func permute(arr []string, l int, r int) {
	if l == r {
		fmt.Println(arr)
		return
	}

	for i := l; i <= r; i++ {
		arr[l], arr[i] = arr[i], arr[l] // Swap
		permute(arr, l+1, r)
		arr[l], arr[i] = arr[i], arr[l] // Backtrack
	}
}

func main() {
	str := []string{"A", "B", "C"}
	permute(str, 0, len(str)-1)
}
