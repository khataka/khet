package main

import "fmt"

func main() {
	var i int
	fmt.Println("Enter an Integer input: ")
	fmt.Scanf("%d", &i)
	f := float64(i)
	fmt.Printf("The float64 representation of %d is %f\n", i, f)
}
