package main

import (
	"fmt"
	"sort"
)

func main() {
	// Input the numbers as a slice
	numbers := []int{1, 2, 3, 4, 5, 6, 2}

	// Use the sort package to reverse the numbers
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))

	// Print the reversed numbers
	fmt.Println("Reversed numbers:", numbers)
}
