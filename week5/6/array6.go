package main

import "fmt"

func main() {
	numbers := [5]int{0, 1, 2, 3, 4}
	fmt.Println(numbers[1])
	numbers[1] = 10
	fmt.Println(numbers[1])
	length := len(numbers)
	fmt.Println("Length = ", length)

}
