package main

import "fmt"

func main() {
	var name string
	var surname string
	var age int
	var height float32
	var weight float32

	fmt.Print("tell me your profile:")
	fmt.Scan(&name, &surname, &age, &height, &weight)
	fmt.Print("i'm: ", name, " ", surname, " ", age, "years old ", height, "cm ", weight, "kg.")

}
