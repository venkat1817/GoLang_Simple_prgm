package main

import "fmt"

func main() {
	limit := 10
	fmt.Println("even numers")
	for i := 0; i < limit; i++ {
		if i%2 == 0 {
			fmt.Println(i)

		}
	}
}
