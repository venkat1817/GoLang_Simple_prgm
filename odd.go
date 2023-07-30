package main

import "fmt"

func main() {
	end := 10

	fmt.Println("Even numbers:")
	for i := 0; i <= end; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	fmt.Println("\nOdd numbers:")
	for i := 0; i <= end; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}
