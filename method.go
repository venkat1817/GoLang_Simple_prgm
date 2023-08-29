// Using method and sum,divition,multiplic two numbers.
package main

import "fmt"

type calculate struct {
	num1, num2 int
}

func (a calculate) sum() int {
	return a.num1 + a.num2
}
func (a calculate) Divi() int {
	return a.num1 - a.num2
}
func (a calculate) mul() int {
	return a.num1 * a.num2
}
func main() {
	aa := calculate{23, 56}
	fmt.Println("sum of two numbers", aa.sum())
	fmt.Println("Divition of two numbers", aa.Divi())
	fmt.Println("multiplication of two numbers", aa.mul())
}
