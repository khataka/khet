package main

import "fmt"

func main() {
	fmt.Print(addTwoNumbers(5, 6))
}

func addTwoNumbers(val1 interface{}, val2 interface{}) int {
	op1, _ := val1.(int)
	op2, _ := val2.(int)

	return op1 + op2
}
