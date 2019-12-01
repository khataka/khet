package main

import "fmt"

func main() {
	var age int
	var height float32
	var weight float32

	fmt.Println("input:")
	fmt.Scanln(&age, &height, &weight)
	fmt.Println("your profile is :", age, "years old,", height, "cm,", weight, "kg.")

}
