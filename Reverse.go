package main

import "fmt"

func reversePrint(s string) {
	runes := []rune(s)
	length := len(runes)

	for i := length - 1; i >= 0; i-- {
		fmt.Println(runes[i])
	}

	fmt.Println()
}

func main() {
	str := "Hello, World!"
	fmt.Println("Original String:", str)
	// fmt.Print("Reversed String: ")
	reversePrint(str)
}
